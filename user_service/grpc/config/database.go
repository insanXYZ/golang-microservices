package config

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var queryMigrate = `
	create table if not exists users (
		id varchar(100) primary key,
		username varchar(50) not null,
		email varchar(50) not null,
		password varchar(255) not null,
		created_at timestamp not null default current_timestamp
	);
`

func migrateTable(ctx context.Context, db *pgx.Conn) error {

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, queryMigrate)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func NewDatabase(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:12345678@db_user:5432/user_service")
	if err != nil {
		return nil, err
	}

	err = migrateTable(ctx, conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
