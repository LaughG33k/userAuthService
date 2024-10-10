package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
)

func (s *Service) Registration(ctx context.Context, user model.User) error {

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
