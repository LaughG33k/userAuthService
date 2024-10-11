package service

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
)

type Auth interface {
	Login(ctx context.Context, login, password string, fp model.FingerPrint) (model.TokenPair, error)
	Logout(context.Context, model.TokenPair) error
	Registration(context.Context, model.User) error
	UpdateSession(context.Context, model.TokenPair, model.FingerPrint) (model.TokenPair, error)
}
