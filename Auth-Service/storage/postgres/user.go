package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "github.com/Time-Capsule/Auth-Service/genproto/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (p *UserRepo) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	var query string
	var args []interface{}
	if req.Email == "" && req.Id == "" {
		return nil, errors.New("either email or id must be provided")
	}
	if req.Email != "" {
		query = `
            SELECT id, username, email, full_name, created_at
            FROM users
            WHERE email = $1
        `
		args = append(args, req.Email)
	} else if req.Id != "" {
		query = `
            SELECT id, username, email, full_name, created_at
            FROM users
            WHERE id = $1
        `
		args = append(args, req.Id)
	} else {
		return nil, errors.New("neither email nor id provided")
	}

	var user pb.GetProfileResponse
	err := p.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt,
	)
	if err != nil {
		log.Printf("Error getting profile: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func (p *UserRepo) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, full_name = $3, updated_at = NOW()
		WHERE id = $4
		`

	_, err := p.db.ExecContext(ctx, query, req.Username, req.Email, req.FullName, req.Id)
	if err != nil {
		log.Printf("Error update profile: %v\n", err)
		return nil, err
	}
	return &pb.UpdateProfileResponse{}, nil

}

func (p *UserRepo) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := p.db.ExecContext(ctx, query, req.NewPassword, req.Id)
	if err != nil {
		log.Printf("Error change password: %v\n", err)
		return nil, err
	}
	return &pb.ChangePasswordResponse{Message: "Password changed successfully"}, nil
}
