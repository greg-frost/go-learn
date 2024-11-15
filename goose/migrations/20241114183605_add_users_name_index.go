package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up, Down)
}

func Up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE INDEX idx_users_name ON users(name);")
	if err != nil {
		return err
	}
	return nil
}

func Down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP INDEX IF EXISTS idx_users_name;")
	if err != nil {
		return err
	}
	return nil
}
