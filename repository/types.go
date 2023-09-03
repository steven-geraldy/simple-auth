// This file contains types that are used in the repository layer.
package repository

type RegisterUserInput struct {
	Name     string
	Phone    string
	Password string
}

type UpdateUserInput struct {
	Name  string
	Phone string
	ID    int
}

type GetUserInput struct {
	ID int
}

type GetUserOutput struct {
	Name  string
	Phone string
}

type LoginInput struct {
	Phone    string
	Password string
}

type LoginOutput struct {
	ID int
}

type GenerateTokenInput struct {
	ID int
}

type GenerateTokenOutput struct {
	Token string
}
type ParseTokenInput struct {
	Token string
}
type ParseTokenOutput struct {
	ID  int
	Exp int64
}
