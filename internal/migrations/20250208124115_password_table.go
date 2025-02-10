package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upPasswordTable, downPasswordTable)
}

func upPasswordTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
		create table if not exists passwords (
		    userid char(14) primary key,
		    password varchar(255) not null,
		    token text not null,
		    created_at int not null,
		    updated_at int not null
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func downPasswordTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table if exists passwords`)
	if err != nil {
		return err
	}
	return nil
}
