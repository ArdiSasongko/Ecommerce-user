package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-user/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-user/internal/model"
	"github.com/ArdiSasongko/Ecommerce-user/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	User interface {
		InsertUser(context.Context, *model.UserPaylod) error
		Login(context.Context, *model.LoginPayload) (*model.LoginResponse, error)
		RefreshToken(context.Context, int32) (*model.LoginResponse, error)
		DeleteSession(context.Context, int32) error
		GetUserByID(context.Context, int32) (*model.UserResponse, error)
	}
	Session interface {
		TokenByToken(context.Context, string) (sqlc.UserSession, error)
	}
}

func NewService(db *pgxpool.Pool, auth auth.JWTAuth) Service {
	q := sqlc.New(db)
	return Service{
		User: &UserService{
			q:    q,
			db:   db,
			auth: auth,
		},
		Session: &SessionService{
			q: q,
		},
	}
}
