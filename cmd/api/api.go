package api

import (
	"github.com/ArdiSasongko/Ecommerce-user/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type Application struct {
	config  Config
	handler handler.Handler
}

type Config struct {
	addrHTTP string
	log      *logrus.Logger
	db       DBConfig
	auth     AuthConfig
}

type DBConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type AuthConfig struct {
	secret string
	iss    string
	aud    string
}

func (a *Application) Mount() *fiber.App {
	r := fiber.New()
	r.Use(recover.New())

	r.Get("/health", a.handler.Health.Check)

	v1 := r.Group("/v1")
	authentication := v1.Group("/authentication")

	authentication.Post("/register/user", a.handler.User.CreateUser)
	authentication.Post("/register/admin", a.handler.User.CreateUser)

	return r
}

func (a *Application) Run(r *fiber.App) error {
	a.config.log.Printf("http server has run, port%v", a.config.addrHTTP)
	return r.Listen(a.config.addrHTTP)
}
