package conf

import (
	"os"

	"go-learn/base"
)

const (
	SITE_NAME      string = "Rest2"
	DEFAULT_LIMIT  int    = 10
	MAX_LIMIT      int    = 1000
	MAX_POST_CHARS int    = 1000
)

var SiteUrl, AbsolutePath string

func init() {
	mode := os.Getenv("MARTINI_ENV")
	path := base.Dir("rest2")

	switch mode {
	case "production":
		SiteUrl = "http://example.com"
		AbsolutePath = path
	default:
		SiteUrl = "http://127.0.0.1"
		AbsolutePath = path
	}
}
