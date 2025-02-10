package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upTnxTable, downTnxTable)
}

func upTnxTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
				create table if not exists tnxs (
				    userId char(14) primary key,
				    amount decimal(20,2) not null,
				    txnTime timestamp default CURRENT_TIMESTAMP,
				    misc JSON
				)
		`)
	if err != nil {
		return err
	}
	return nil
}

func downTnxTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table if exists tnxs`)
	if err != nil {
		return err
	}
	return nil
}
