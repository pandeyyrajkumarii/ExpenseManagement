package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUsersTable, downUsersTable)
}

func upUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		create table if not exists users (
			id char(14) NOT NULL primary key,
			name varchar(50) NOT NULL,
			age int NOT NULL,
			gender char(1) NOT NULL, 
			created_at int NOT NULL,
			updated_at int NOT NULL);
		`)
	if err != nil {
		return err
	}
	return nil
}

func downUsersTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table if exists users;`)
	if err != nil {
		return err
	}
	return nil
}
