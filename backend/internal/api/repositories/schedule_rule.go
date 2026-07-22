package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/api/models"
)

type ScheduleRuleRepository interface {
	// GetByUser возвращает актуальные правила расписания пользователя (valid_to IS NULL).
	GetByUser(ctx context.Context, userID int) ([]models.ScheduleRule, error)
	GetByID(ctx context.Context, id int) (*models.ScheduleRule, error)
	Create(ctx context.Context, rule models.ScheduleRule) (*models.ScheduleRule, error)
	// Replace "закрывает" старое правило (valid_to = now) и создаёт новое — без потери истории.
	Replace(ctx context.Context, oldID int, rule models.ScheduleRule) (*models.ScheduleRule, error)
	// GetGroups возвращает id групп, привязанных к правилу.
	GetGroups(ctx context.Context, ruleID int) ([]int, error)
	SetGroups(ctx context.Context, ruleID int, groupIDs []int) error
}

type scheduleRuleRepository struct {
	db *pgxpool.Pool
}

func NewScheduleRuleRepository(db *pgxpool.Pool) ScheduleRuleRepository {
	return &scheduleRuleRepository{db: db}
}

func (r *scheduleRuleRepository) GetByUser(ctx context.Context, userID int) ([]models.ScheduleRule, error) {
	const query = `
		SELECT id, user_id, work_type_id, subject_name, day_of_week,
		       week_parity, start_time, end_time, room, valid_from, valid_to
		FROM schedule_rules
		WHERE user_id = $1 AND valid_to IS NULL
		ORDER BY day_of_week, start_time`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.GetByUser: %w", err)
	}
	defer rows.Close()

	var rules []models.ScheduleRule
	for rows.Next() {
		rule, err := scanScheduleRule(rows)
		if err != nil {
			return nil, fmt.Errorf("scheduleRuleRepository.GetByUser scan: %w", err)
		}
		rules = append(rules, *rule)
	}
	return rules, rows.Err()
}

func (r *scheduleRuleRepository) GetByID(ctx context.Context, id int) (*models.ScheduleRule, error) {
	const query = `
		SELECT id, user_id, work_type_id, subject_name, day_of_week,
		       week_parity, start_time, end_time, room, valid_from, valid_to
		FROM schedule_rules
		WHERE id = $1`

	rule, err := scanScheduleRule(r.db.QueryRow(ctx, query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.GetByID: %w", err)
	}
	return rule, nil
}

func (r *scheduleRuleRepository) Create(ctx context.Context, rule models.ScheduleRule) (*models.ScheduleRule, error) {
	const query = `
		INSERT INTO schedule_rules
		    (user_id, work_type_id, subject_name, day_of_week, week_parity,
		     start_time, end_time, room, valid_from)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, user_id, work_type_id, subject_name, day_of_week,
		          week_parity, start_time, end_time, room, valid_from, valid_to`

	created, err := scanScheduleRule(r.db.QueryRow(ctx, query,
		rule.UserID, rule.WorkTypeID, rule.SubjectName, rule.DayOfWeek,
		rule.WeekParity, rule.StartTime, rule.EndTime, rule.Room, rule.ValidFrom,
	))
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.Create: %w", err)
	}
	return created, nil
}

func (r *scheduleRuleRepository) Replace(ctx context.Context, oldID int, rule models.ScheduleRule) (*models.ScheduleRule, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.Replace begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	// Закрываем старое правило
	now := time.Now()
	tag, err := tx.Exec(ctx,
		`UPDATE schedule_rules SET valid_to = $2 WHERE id = $1 AND valid_to IS NULL`,
		oldID, now,
	)
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.Replace close old: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return nil, ErrNotFound
	}

	// Создаём новое правило
	const insertQuery = `
		INSERT INTO schedule_rules
		    (user_id, work_type_id, subject_name, day_of_week, week_parity,
		     start_time, end_time, room, valid_from)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, user_id, work_type_id, subject_name, day_of_week,
		          week_parity, start_time, end_time, room, valid_from, valid_to`

	rule.ValidFrom = now
	created, err := scanScheduleRule(tx.QueryRow(ctx, insertQuery,
		rule.UserID, rule.WorkTypeID, rule.SubjectName, rule.DayOfWeek,
		rule.WeekParity, rule.StartTime, rule.EndTime, rule.Room, rule.ValidFrom,
	))
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.Replace create new: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.Replace commit: %w", err)
	}
	return created, nil
}

func (r *scheduleRuleRepository) GetGroups(ctx context.Context, ruleID int) ([]int, error) {
	const query = `SELECT group_id FROM schedule_rule_groups WHERE schedule_rule_id = $1`

	rows, err := r.db.Query(ctx, query, ruleID)
	if err != nil {
		return nil, fmt.Errorf("scheduleRuleRepository.GetGroups: %w", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scheduleRuleRepository.GetGroups scan: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

func (r *scheduleRuleRepository) SetGroups(ctx context.Context, ruleID int, groupIDs []int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("scheduleRuleRepository.SetGroups begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`DELETE FROM schedule_rule_groups WHERE schedule_rule_id = $1`, ruleID,
	); err != nil {
		return fmt.Errorf("scheduleRuleRepository.SetGroups delete: %w", err)
	}

	for _, gid := range groupIDs {
		if _, err := tx.Exec(ctx,
			`INSERT INTO schedule_rule_groups (schedule_rule_id, group_id) VALUES ($1, $2)`,
			ruleID, gid,
		); err != nil {
			return fmt.Errorf("scheduleRuleRepository.SetGroups insert group %d: %w", gid, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("scheduleRuleRepository.SetGroups commit: %w", err)
	}
	return nil
}

// scanScheduleRule принимает pgx.Row или pgx.Rows — оба удовлетворяют интерфейсу pgx.Row.
func scanScheduleRule(row pgx.Row) (*models.ScheduleRule, error) {
	var r models.ScheduleRule
	err := row.Scan(
		&r.ID, &r.UserID, &r.WorkTypeID, &r.SubjectName, &r.DayOfWeek,
		&r.WeekParity, &r.StartTime, &r.EndTime, &r.Room, &r.ValidFrom, &r.ValidTo,
	)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
