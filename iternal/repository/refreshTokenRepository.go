package repository

import (
	"context"

	"github.com/LaughG33k/userAuthService/iternal/dbClient/postgresql"
)

type RefreshTokenRepository struct {
	client postgresql.Client
	ctx    context.Context
}

func NewRefreshTokenRepostiroy(ctx context.Context, client postgresql.Client) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *RefreshTokenRepository) CreateRefreshToken(oldRefreshToken, refreshToken, ownerUuid string, tokenTimeLife int64) error {

	if oldRefreshToken != "" {

		tx, err := r.client.Begin(r.ctx)

		if err != nil {
			return err
		}

		defer tx.Rollback(r.ctx)

		if _, err := tx.Exec(r.ctx, "delete from refresh_tokens where token=$1 and owner_uuid=$2;", oldRefreshToken, ownerUuid); err != nil {
			return err
		}

		if _, err := tx.Exec(r.ctx, "insert into refresh_tokens(token, owner_uuid, time_end_of_life) values($1, $2, $3);", refreshToken, ownerUuid, tokenTimeLife); err != nil {
			return err
		}

		if err := tx.Commit(r.ctx); err != nil {
			return err
		}

		return nil
	}

	if _, err := r.client.Exec(r.ctx, "insert into refresh_tokens(token, owner_uuid, time_end_of_life) values($1, $2, $3);", refreshToken, ownerUuid, tokenTimeLife); err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepository) FindRefreshToken(refreshToken string) (string, int64, error) {

	ownerUuid := ""
	var timeLife int64 = 0

	if err := r.client.QueryRow(r.ctx, "select owner_uuid, time_end_of_life from refresh_tokens where token=$1;", refreshToken).Scan(&ownerUuid, &timeLife); err != nil {
		return "", 0, err
	}

	return ownerUuid, timeLife, nil

}
