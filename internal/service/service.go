package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-user/internal/model"
	"github.com/ArdiSasongko/Ecommerce-user/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	User interface {
		InsertUser(context.Context, *model.UserPaylod) error
	}
}

func NewService(db *pgxpool.Pool) Service {
	q := sqlc.New(db)
	return Service{
		User: &UserService{
			q:  q,
			db: db,
		},
	}
}
