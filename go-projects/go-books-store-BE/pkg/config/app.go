package config

import (
	_ "github.com/lib/pq"
)

type ConfigType string

const (
	host     ConfigType = "localhost"
	port     int        = 5432
	user     ConfigType = "postgres"
	password ConfigType = "1"
	dbname   ConfigType = "bookstore"
)
