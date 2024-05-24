package repository

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/dbClient/postgresql"
)

type RefreshTokenRepository struct {
	client postgresql.Client
}

func NewRefreshTokenRepostiroy(client postgresql.Client) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		client: client,
	}
}

func (r *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, oldRefreshToken, refreshToken, ownerUuid string, tokenTimeLife int64) error {

	if oldRefreshToken != "" {

		tx, err := r.client.Begin(ctx)

		if err != nil {
			return err
		}

		defer tx.Rollback(ctx)

		if _, err := tx.Exec(ctx, "delete from refresh_tokens where token=$1 and owner_uuid=$2;", oldRefreshToken, ownerUuid); err != nil {
			return err
		}

		if _, err := tx.Exec(ctx, "insert into refresh_tokens(token, owner_uuid, time_end_of_life) values($1, $2, $3);", refreshToken, ownerUuid, tokenTimeLife); err != nil {
			return err
		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}

		return nil
	}

	if _, err := r.client.Exec(ctx, "insert into refresh_tokens(token, owner_uuid, time_end_of_life) values($1, $2, $3);", refreshToken, ownerUuid, tokenTimeLife); err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepository) FindRefreshToken(ctx context.Context, refreshToken string) (string, int64, error) {

	ownerUuid := ""
	var timeLife int64 = 0

	if err := r.client.QueryRow(ctx, "select owner_uuid, time_end_of_life from refresh_tokens where token=$1;", refreshToken).Scan(&ownerUuid, &timeLife); err != nil {
		return "", 0, err
	}

	return ownerUuid, timeLife, nil

}
