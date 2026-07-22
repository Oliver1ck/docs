package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Oliver1ck/docs/internal/api/models"
	"github.com/Oliver1ck/docs/internal/api/repositories"
	"github.com/Oliver1ck/docs/internal/config"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*models.User, error)
	Login(ctx context.Context, email, password string) (*TokenPair, error)
	GetByID(ctx context.Context, id int) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
	jwt  config.JWT
}

func NewUserService(repo repositories.UserRepository, jwtCfg config.JWT) UserService {
	return &userService{repo: repo, jwt: jwtCfg}
}

func (s *userService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	_, err := s.repo.GetByEmail(ctx, email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	if !errors.Is(err, repositories.ErrNotFound) {
		return nil, fmt.Errorf("userService.Register check email: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("userService.Register hash password: %w", err)
	}

	user, err := s.repo.Create(ctx, models.User{
		Username: username,
		Email:    email,
		Password: string(hash),
		RoleID:   1,
	})
	if err != nil {
		return nil, fmt.Errorf("userService.Register create: %w", err)
	}
	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (*TokenPair, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if errors.Is(err, repositories.ErrNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, fmt.Errorf("userService.Login get user: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	tokens, err := s.generateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("userService.Login generate tokens: %w", err)
	}
	return tokens, nil
}

func (s *userService) GetByID(ctx context.Context, id int) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetByID: %w", err)
	}
	return user, nil
}

func (s *userService) generateTokenPair(user *models.User) (*TokenPair, error) {
	now := time.Now()

	accessClaims := jwt.MapClaims{
		"sub":     user.ID,
		"role_id": user.RoleID,
		"exp":     now.Add(s.jwt.AccessTokenTTL).Unix(),
		"iat":     now.Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).
		SignedString([]byte(s.jwt.Secret))
	if err != nil {
		return nil, fmt.Errorf("sign access token: %w", err)
	}

	refreshClaims := jwt.MapClaims{
		"sub": user.ID,
		"exp": now.Add(s.jwt.RefreshTokenTTL).Unix(),
		"iat": now.Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(s.jwt.Secret))
	if err != nil {
		return nil, fmt.Errorf("sign refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
