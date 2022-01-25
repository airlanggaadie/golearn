package common

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

// Pool Define pgx pool connection database
var Pool *pgxpool.Pool

// Init /** initialize database connection with database url
func Init() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), ViperEnvVariable(new(ViperParameters).KeyWithDefaultConfig("DATABASE_URL", "")))
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("%#T %#v", err, err)
	}

	Pool = pool
	return Pool
}

// GetDB /** return database connection instance
func GetDB() *pgxpool.Pool {
	return Pool
}