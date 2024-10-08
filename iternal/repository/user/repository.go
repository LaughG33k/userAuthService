package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/LaughG33k/userAuthService/client/postgresql"
	"github.com/LaughG33k/userAuthService/iternal/model"
	"github.com/LaughG33k/userAuthService/iternal/repository"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	client postgresql.Client
}

func NewUserRepostiroy(client postgresql.Client) repository.UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Create(ctx context.Context, user model.User) error {

	if _, err := r.client.Exec(ctx, "insert into users(login, password, name, email) values($1, $2, $3, $4) returning uuid;", user.Login, user.Password, user.Name, user.Email); err != nil {

		return err

	}

	return nil
}

func (r *UserRepository) GetUuidByLP(ctx context.Context, user model.User) (string, error) {

	uuid := ""

	if err := r.client.QueryRow(ctx, "select uuid from users where login=$1 and password=$2;", user.Login, user.Password).Scan(&uuid); err != nil {

		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			return "", fmt.Errorf(pgError.Code)
		}

		return "", err

	}

	return uuid, nil
}
