package service

import (
	"context"
	"errors"
	"time"

	"github.com/ArdiSasongko/Ecommerce-user/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-user/internal/model"
	"github.com/ArdiSasongko/Ecommerce-user/internal/storage/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

const layout = "2006-01-02"

var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicatePhone    = errors.New("phone number already exists")
)

type UserService struct {
	q    *sqlc.Queries
	db   *pgxpool.Pool
	auth auth.JWTAuth
}

func (s *UserService) InsertUser(ctx context.Context, payload *model.UserPaylod) error {
	parseDate, err := time.Parse(layout, payload.DoB)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	role, err := qtx.GetRoleByName(ctx, payload.Role)
	if err != nil {
		return err
	}

	_, err = qtx.InsertUser(ctx, sqlc.InsertUserParams{
		Username:    payload.Username,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Address:     payload.Address,
		Dob: pgtype.Date{
			Time:  parseDate,
			Valid: true,
		},
		Password: string(password),
		Fullname: payload.Fullname,
		Role:     role.Level,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				switch pgErr.ConstraintName {
				case "users_email_key":
					return ErrDuplicateEmail
				case "users_username_key":
					return ErrDuplicateUsername
				case "users_phone_number_key":
					return ErrDuplicatePhone
				}
			}
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *UserService) getUserEmail(ctx context.Context, email string) (sqlc.GetUserByEmailRow, error) {
	user, err := s.q.GetUserByEmail(ctx, email)
	if err != nil {
		return sqlc.GetUserByEmailRow{}, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, payload *model.LoginPayload) (*model.LoginResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := s.q.WithTx(tx)

	user, err := s.getUserEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, err
	}

	active_token, err := s.auth.GeneratedToken(user.ID, "active_token")
	if err != nil {
		return nil, err
	}

	refresh_token, err := s.auth.GeneratedToken(user.ID, "refresh_token")
	if err != nil {
		return nil, err
	}

	sessionPayload := sqlc.InsertSessionParams{
		UserID:       user.ID,
		ActiveToken:  active_token,
		RefreshToken: refresh_token,
		ActiveTokenExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(auth.TokenTime["active_token"]),
			Valid: true,
		},
		RefreshTokenExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(auth.TokenTime["refresh_token"]),
			Valid: true,
		},
	}

	_, err = qtx.InsertSession(ctx, sessionPayload)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		ActiveToken:  active_token,
		RefreshToken: refresh_token,
	}, nil
}
