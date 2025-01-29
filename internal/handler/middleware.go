package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArdiSasongko/Ecommerce-user/internal/config/auth"
	"github.com/ArdiSasongko/Ecommerce-user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareHandler struct {
	service service.Service
	auth    auth.JWTAuth
}

func (h *MiddlewareHandler) TokenMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == " " {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing token header authorization",
			})
		}

		rContext := ctx.Context()
		parts := strings.Split(authToken, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "header are malformed",
			})
		}

		token := parts[1]

		// check if token exists
		_, err := h.service.Session.TokenByToken(rContext, token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "active token invalid",
			})
		}

		jwttoken, err := h.auth.ValidateRefreshToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		claims, _ := jwttoken.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		ctx.Locals("user_id", userID)
		return ctx.Next()
	}
}
