package aar

import (
	"github.com/jackc/pgx"
)

var (
	// DB is the pool for database connections
	DB *pgx.ConnPool
)
