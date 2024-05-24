package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LaughG33k/userAuthService/iternal/dbClient/postgresql"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	client postgresql.Client
}

func NewUserRepostiroy(client postgresql.Client) *UserRepository {
	return &UserRepository{

		client: client,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, login, password, name, email string) error {

	if _, err := r.client.Exec(ctx, "insert into users(login, password, name, email) values($1, $2, $3, $4);", login, password, name, email); err != nil {

		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			return fmt.Errorf(pgError.Code)
		}

		return err

	}

	return nil
}

func (r *UserRepository) CheckUserByLP(ctx context.Context, login, password string) (string, error) {

	uuid := ""

	if err := r.client.QueryRow(ctx, "select uuid from users where login=$1 and password=$2;", login, password).Scan(&uuid); err != nil {

		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			return "", fmt.Errorf(pgError.Code)
		}

		return "", err

	}

	return uuid, nil
}
