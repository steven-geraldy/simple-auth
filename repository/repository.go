// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

type NewUserRepositoryOptions struct {
	Dsn string
}

func NewUserRepository(opts NewUserRepositoryOptions) *UserRepository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		db: db,
	}
}

type JWTRepository struct {
	expiryTime time.Duration
	privateKey string
	publicKey  string
}

type NewJWTRepositoryOptions struct {
	ExpiryTime time.Duration
	PrivateKey string
	PublicKey  string
}

func NewJWTRepository(opts NewJWTRepositoryOptions) *JWTRepository {
	return &JWTRepository{
		expiryTime: opts.ExpiryTime,
		privateKey: opts.PrivateKey,
		publicKey:  opts.PublicKey,
	}
}
