package sqlstore_test

import (
	"os"
	"testing"
)

// Путь подключения к БД
var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=postgres password=admin dbname=learn_test sslmode=disable"
	}

	os.Exit(m.Run())
}
