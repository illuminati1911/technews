package env

import (
	"os"
)

// TODO: move to .local.env file
var localDefaults = map[string]string{
	"POSTGRES_USER":     "postgres",
	"POSTGRES_PASSWORD": "password",
	"POSTGRES_DB":       "technews",
	"POSTGRES_HOST":     "localhost",
	"POSTGRES_PORT":     "5432",
	"POSTGRES_RETRIES":  "5",
	"JWT_SECRET":        "SECRET",
}

// Init will check that above listed env vars exists
// or it will load the defaults for local development.
func init() {
	for k, v := range localDefaults {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
