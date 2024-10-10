package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
)

func (s *Service) Logout(ctx context.Context, token model.TokenPair) error {

	if _, err := s.jwtParse.ParseToken(token.Jwt); err != nil {
		return err
	}

	if err := s.sessionRepo.Delete(ctx, token.Refresh); err != nil {
		return err
	}

	return nil
}
