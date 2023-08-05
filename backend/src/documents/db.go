package documents

import (
	"context"
	"database/sql"
	"sapp/paperless-accounting/database"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

var ddl string

func (d *DocumentMgr) setupDB() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		return err
	}

	// create tables if necessary
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	d.db = database.New(db)

	return nil
}
