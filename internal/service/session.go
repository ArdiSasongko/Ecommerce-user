package service

import (
	"context"

	"github.com/ArdiSasongko/Ecommerce-user/internal/storage/sqlc"
)

type SessionService struct {
	q *sqlc.Queries
}

func (s *SessionService) TokenByToken(ctx context.Context, active_token string) (sqlc.UserSession, error) {
	token, err := s.q.GetSessionByActiveToken(ctx, active_token)
	if err != nil {
		return sqlc.UserSession{}, err
	}

	return token, nil
}
