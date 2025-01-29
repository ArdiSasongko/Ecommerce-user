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

func (h *MiddlewareHandler) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == "" {
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

		// check if token exists in session
		_, err := h.service.Session.TokenByToken(rContext, token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "active token invalid",
			})
		}

		// validate JWT token
		jwttoken, err := h.auth.ValidateActiveToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// validate claims
		claims, ok := jwttoken.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		// check if subject exists in claims
		sub, exists := claims["sub"]
		if !exists {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing subject claim",
			})
		}

		// parse user ID
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", sub), 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user ID format",
			})
		}

		ctx.Locals("user_id", userID)
		return ctx.Next()
	}
}

func (h *MiddlewareHandler) TokenMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == "" {
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

		// validate JWT token
		jwttoken, err := h.auth.ValidateRefreshToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// validate claims
		claims, ok := jwttoken.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		// check if subject exists in claims
		sub, exists := claims["sub"]
		if !exists {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing subject claim",
			})
		}

		// parse user ID
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", sub), 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user ID format",
			})
		}

		ctx.Locals("user_id", userID)
		return ctx.Next()
	}
}
