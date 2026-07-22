-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles (
    id   SERIAL      PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE users (
    id         SERIAL       PRIMARY KEY,
    username   VARCHAR(100) NOT NULL UNIQUE,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    role_id    INTEGER      NOT NULL REFERENCES roles(id),
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE permissions (
    id   SERIAL       PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE user_permissions (
    user_id       INTEGER NOT NULL REFERENCES users(id)       ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

CREATE TABLE groups (
    id         SERIAL       PRIMARY KEY,
    name       VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE work_categories (
    id         SERIAL       PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    sort_order INTEGER      NOT NULL DEFAULT 0
);

CREATE TABLE work_types (
    id          SERIAL       PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    short_name  VARCHAR(50)  NOT NULL,
    category_id INTEGER      NOT NULL REFERENCES work_categories(id),
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE
);

CREATE TABLE schedule_rules (
    id           SERIAL       PRIMARY KEY,
    user_id      INTEGER      NOT NULL REFERENCES users(id)      ON DELETE CASCADE,
    work_type_id INTEGER      NOT NULL REFERENCES work_types(id),
    subject_name VARCHAR(255) NOT NULL,
    day_of_week  SMALLINT     NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
    week_parity  SMALLINT     NOT NULL DEFAULT 0 CHECK (week_parity IN (0, 1, 2)),
    start_time   TIME         NOT NULL,
    end_time     TIME         NOT NULL,
    room         VARCHAR(100) NOT NULL DEFAULT '',
    valid_from   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    valid_to     TIMESTAMPTZ
);

CREATE TABLE schedule_rule_groups (
    schedule_rule_id INTEGER NOT NULL REFERENCES schedule_rules(id) ON DELETE CASCADE,
    group_id         INTEGER NOT NULL REFERENCES groups(id)         ON DELETE CASCADE,
    PRIMARY KEY (schedule_rule_id, group_id)
);

CREATE TABLE tracks (
    id               SERIAL         PRIMARY KEY,
    user_id          INTEGER        NOT NULL REFERENCES users(id)          ON DELETE CASCADE,
    schedule_rule_id INTEGER                 REFERENCES schedule_rules(id) ON DELETE SET NULL,
    work_type_id     INTEGER        NOT NULL REFERENCES work_types(id),
    date             DATE           NOT NULL,
    academic_hours   NUMERIC(5, 1)  NOT NULL CHECK (academic_hours > 0),
    status           VARCHAR(50)    NOT NULL DEFAULT 'pending',
    comment          TEXT           NOT NULL DEFAULT ''
);

CREATE TABLE track_groups (
    track_id INTEGER NOT NULL REFERENCES tracks(id) ON DELETE CASCADE,
    group_id INTEGER NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (track_id, group_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS track_groups;
DROP TABLE IF EXISTS tracks;
DROP TABLE IF EXISTS schedule_rule_groups;
DROP TABLE IF EXISTS schedule_rules;
DROP TABLE IF EXISTS work_types;
DROP TABLE IF EXISTS work_categories;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS user_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

-- +goose StatementEnd
