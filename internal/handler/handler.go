package handler

import (
	"github.com/ArdiSasongko/Ecommerce-user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Health interface {
		Check(*fiber.Ctx) error
	}
	User interface {
		CreateUser(ctx *fiber.Ctx) error
	}
}

func NewHandler(db *pgxpool.Pool) Handler {
	service := service.NewService(db)
	return Handler{
		Health: &HealthHandler{},
		User: &UserHandler{
			service: service,
		},
	}
}
