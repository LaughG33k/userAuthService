package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/LaughG33k/userAuthService/iternal/dbClient/postgresql"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	ctx    context.Context
	client postgresql.Client
}

func NewUserRepostiroy(ctx context.Context, client postgresql.Client) *UserRepository {
	return &UserRepository{

		ctx:    ctx,
		client: client,
	}
}

func (r *UserRepository) CreateUser(login, password, name, email string) error {

	if _, err := r.client.Exec(r.ctx, "insert into users(login, password, name, email) values($1, $2, $3, $4);", login, password, name, email); err != nil {

		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			return fmt.Errorf(pgError.Code)
		}

		return err

	}

	return nil
}

func (r *UserRepository) CheckUserByLP(login, password string) (string, error) {

	uuid := ""

	if err := r.client.QueryRow(r.ctx, "select uuid from users where login=$1 and password=$2;", login, password).Scan(&uuid); err != nil {

		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			return "", fmt.Errorf(pgError.Code)
		}

		return "", err

	}

	return uuid, nil
}
