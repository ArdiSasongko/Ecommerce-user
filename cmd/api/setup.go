package api

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-user/internal/config/env"
	"github.com/ArdiSasongko/Ecommerce-user/internal/config/logger"
	"github.com/ArdiSasongko/Ecommerce-user/internal/config/pg"
	"github.com/ArdiSasongko/Ecommerce-user/internal/handler"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	logger := logger.NewLogger()
	config := Config{
		addrHTTP: env.GetEnvString("ADDR_HTTP", ""),
		log:      logger,
		auth: AuthConfig{
			secret: env.GetEnvString("JWT_SECRET", ""),
			iss:    env.GetEnvString("JWT_ISS", ""),
			aud:    env.GetEnvString("JWT_AUD", ""),
		},
	}

	return config, nil
}

func ConnDatabase(cfg DBConfig, log *logrus.Logger) (*pgxpool.Pool, error) {
	conn, err := pg.New(cfg.addr, cfg.maxOpenConns, cfg.maxIdleConns, cfg.maxIdleTime)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Info("Succes Connected to database")
	return conn, nil
}

func SetupHTTPApplication() (*Application, error) {
	cfg, err := LoadConfig()
	if err != nil {
		cfg.log.Fatalf("%s", err.Error())
	}

	// _, err = ConnDatabase(cfg.db, cfg.log)
	// if err != nil {
	// 	cfg.log.Fatalf("failed to connected database :%v", err)
	// }

	//auth := auth.NewJWT(cfg.auth.secret, cfg.auth.aud, cfg.auth.iss)
	handler := handler.NewHandler()
	return &Application{
		config:  cfg,
		handler: handler,
	}, nil
}
