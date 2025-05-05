package config

import (
	"profkom/pkg/postgres"
	"profkom/pkg/s3"
)

type Config struct {
	Postgres postgres.Config
	S3       s3.Config
}
