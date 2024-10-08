package auth

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
)

func (a *Auth) Registration(ctx context.Context, user model.User) error {

	if err := a.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
