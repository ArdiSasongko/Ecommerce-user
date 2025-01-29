package handler

import (
	"github.com/ArdiSasongko/Ecommerce-user/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
	User interface {
		Register(*fiber.Ctx) error
		Login(*fiber.Ctx) error
	}
}

func NewHandler(db *pgxpool.Pool, auth auth.JWTAuth) Handler {
	service := service.NewService(db, auth)
	return Handler{
		Health: &HealthHandler{},
		User: &UserHandler{
			service: service,
		},
	}
}
