package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/LaughG33k/userAuthService/iternal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

func NewClient(ctx context.Context, trysToConnNum int, cfg config.DBConfig) (pool *pgxpool.Pool, err error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	err = trysToConnect(func() error {

		timeCtx, cancle := context.WithTimeout(ctx, 5*time.Second)
		defer cancle()

		pool, err = pgxpool.New(timeCtx, connString)

		if err != nil {
			return err
		}

		if err = pool.Ping(timeCtx); err != nil {
			return err
		}

		return nil

	}, trysToConnNum, 3*time.Second)

	if err != nil {
		return nil, err
	}

	return pool, nil

}

func trysToConnect(fn func() error, trysToConnNum int, timing time.Duration) (err error) {

	for trysToConnNum > 0 {

		if err = fn(); err != nil {

			time.Sleep(timing)
			trysToConnNum--

			continue
		}

		return nil

	}

	return err

}
