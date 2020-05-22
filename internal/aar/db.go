package aar

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// DB is the pool for database connections
	DB *pgxpool.Pool
)
