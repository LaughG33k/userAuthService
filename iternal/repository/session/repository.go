package session

import (
	"context"

	"github.com/LaughG33k/userAuthService/client/postgresql"
	"github.com/LaughG33k/userAuthService/iternal/model"
	"github.com/LaughG33k/userAuthService/iternal/repository"
)

type SessionRepository struct {
	client postgresql.Client
}

func (r *SessionRepository) Delete(ctx context.Context, token string) error {

	q := "delete from sessions where token = $1;"

	if _, err := r.client.Exec(ctx, q, token); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Create(ctx context.Context, session model.Session) error {

	q := "insert into sessions(token, owner, life_time, addr, browser, device) values ($1, $2, $3, $4, $5, $6);"

	if _, err := r.client.Exec(ctx, q, session.Token, session.Owner, session.LifeTime, session.Addr, session.Browser, session.Device); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Get(ctx context.Context, token string) (model.Session, error) {

	var session model.Session
	q := "select owner, life_time, addr, browser, device from sessions where token = $1;"

	if err := r.client.QueryRow(ctx, q, token).Scan(&session.Owner, &session.LifeTime, &session.Addr, &session.Browser, &session.Device); err != nil {
		return session, err
	}

	return session, nil
}

func NewRefreshTokenRepostiroy(client postgresql.Client) repository.SessionRepository {
	return &SessionRepository{
		client: client,
	}
}
