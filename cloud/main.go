package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go-learn/base"
)

// Путь
var path = base.Dir("cloud")

func init() {
	os.Setenv("TLOG_TYPE", "file")

	os.Setenv("TLOG_FILE", filepath.Join(
		path, "data", "transaction.log"))

	os.Setenv("TLOG_DB_NAME", "learn")
	os.Setenv("TLOG_DB_HOST", "localhost")
	os.Setenv("TLOG_DB_USER", "postgres")
	os.Setenv("TLOG_DB_PASS", "admin")

	os.Setenv("FRONTEND_TYPE", "rest")
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")
}
