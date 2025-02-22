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
				    txnId varchar(255) not null primary key,
				    userId char(14) NOT NULL,
				    amount decimal(20,2) not null,
				    category varchar(255) ,
				    transactionType enum('debit', 'credit') NOT NULL ,
				    description text ,
				    txnTime timestamp default CURRENT_TIMESTAMP,
				    created_at int NOT NULL,
					updated_at int NOT NULL,
				    Foreign key (userId) references users(id) on delete cascade
				);
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
