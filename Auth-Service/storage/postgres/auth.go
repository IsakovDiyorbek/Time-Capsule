package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	pb "github.com/Time-Capsule/Auth-Service/genproto/auth"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type AuthRepo struct {
	db    *sql.DB
	redis *redis.Client
}

func NewAuthRepo(db *sql.DB, redis *redis.Client) *AuthRepo {
	return &AuthRepo{db: db, redis: redis}
}

func (p *AuthRepo) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userId := uuid.NewString()
	query := `
		INSERT INTO users(id, username, email, password_hash, full_name, date_of_birth)
		VALUES ($1, $2, $3, $4, $5, $6)
		`
	_, err := p.db.ExecContext(ctx, query, userId, req.Username, req.Email, req.Password, req.FullName, req.DateOfBirth)

	if err != nil {
		return nil, err
	}

	token := req.Token
	if p.redis == nil {
		return nil, errors.New("redis client is not initialized")
	}

	err = p.redis.Set(ctx, req.Username, token, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	err = p.redis.Set(ctx, token, req.Username, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{}, nil
}

func (p *AuthRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	query := `
		SELECT username
		FROM users
		WHERE username = $1 AND password_hash = $2
	`
	var username string
	err := p.db.QueryRowContext(ctx, query, req.Username, req.Password).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found or invalid credentials
		}
		return nil, err
	}

	token, err := p.redis.Get(ctx, username).Result()
	if err != nil {
		if err == redis.Nil {
			slog.Info(err.Error())
			return nil, nil // Token not found in Redis
		}
		return nil, err
	}

	return &pb.LoginResponse{
		Token:     token,
		ExpiresAt: fmt.Sprintf("%d", time.Now().Add(30*time.Minute).Unix()),
	}, nil
}

func (p *AuthRepo) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	username, err := p.redis.Get(ctx, req.Token).Result()
	if err != nil {
		if err == redis.Nil {
			return &pb.LogoutResponse{Message: "Token not found"}, nil
		}
		return nil, err
	}

	err = p.redis.Del(ctx, username, req.Token).Err()
	if err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{Message: "Logged out successfully"}, nil
}

func (p *AuthRepo) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1
		WHERE email = $2 AND username = $3
	`
	_, err := p.db.ExecContext(ctx, query, req.NewPassword, req.Email, req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.ResetPasswordResponse{Message: "Password reset successfully"}, nil
}
