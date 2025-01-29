package handler

import (
	"strings"

	"github.com/ArdiSasongko/Ecommerce-user/internal/config/logger"
	"github.com/ArdiSasongko/Ecommerce-user/internal/model"
	"github.com/ArdiSasongko/Ecommerce-user/internal/service"
	"github.com/ArdiSasongko/Ecommerce-user/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var log = logger.NewLogger()

type UserHandler struct {
	service service.Service
}

func (s *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	payload := new(model.UserPaylod)

	if err := ctx.BodyParser(payload); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("Body Parser :%w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if strings.Contains(ctx.Route().Path, "user") {
		payload.Role = "user"
	}
	if strings.Contains(ctx.Route().Path, "admin") {
		payload.Role = "admin"
	}

	if err := payload.Validate(); err != nil {
		log.WithError(fiber.ErrBadRequest).Error("validate error :%w", err)
		errs := utils.ValidationError(err.(validator.ValidationErrors))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errs,
		})
	}

	if err := s.service.User.InsertUser(ctx.Context(), payload); err != nil {
		log.WithError(fiber.ErrInternalServerError).Error("error :%w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "ok",
	})
}
