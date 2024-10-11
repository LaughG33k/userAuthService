package repository

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/model"
)

type UserRepository interface {
	Create(context.Context, model.User) error
	GetUuidByLP(ctx context.Context, login, password string) (string, error)
}

type SessionRepository interface {
	Create(context.Context, model.Session) error
	Get(ctx context.Context, token string) (model.Session, error)
	Delete(ctx context.Context, token string) error
}
