// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type UserInterface interface {
	Register(ctx context.Context, input RegisterUserInput) (err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (err error)
	GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error)
	Login(ctx context.Context, input LoginInput) (output LoginOutput, err error)
}

type AuthInterface interface {
	GenerateToken(input GenerateTokenInput) (output GenerateTokenOutput, err error)
	ParseToken(input ParseTokenInput) (output ParseTokenOutput, err error)
}
