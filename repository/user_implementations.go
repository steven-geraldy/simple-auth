package repository

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

const (
	passwordSalt string = "2UuhANkm1g45y6M7u6MCCw=="
)

var _ UserInterface = &UserRepository{}

func (r *UserRepository) Register(ctx context.Context, input RegisterUserInput) (err error) {
	hashedPassword := r.hashAndSalt(ctx, input.Password)
	_, err = r.db.ExecContext(ctx, "INSERT INTO users (name, phone, password) VALUES ($1, $2, $3);",
		input.Name, input.Phone, hashedPassword)
	if err != nil {
		log.Error("[Repository:InsertUser] Failed to insert user data, error"+err.Error(), " | context_id: ", ctx)
		return
	}
	return
}

func (r *UserRepository) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	_, err = r.db.ExecContext(ctx, "UPDATE users set name=$1, phone=$2 WHERE id = $3;",
		input.Name, input.Phone, input.ID)
	if err != nil {
		log.Error("[Repository:UpdateUser] Failed to update user data, error"+err.Error(), " | context_id: ", ctx)
		return
	}
	return
}

func (r *UserRepository) GetUser(ctx context.Context, input GetUserInput) (output GetUserOutput, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT name, phone from users WHERE id = $1;",
		input.ID).Scan(&output.Name, &output.Phone)
	if err != nil {
		log.Error("[Repository:GetUser] Failed to get user data, error"+err.Error(), " | context_id: ", ctx)
		return
	}
	return
}

func (r *UserRepository) Login(ctx context.Context, input LoginInput) (output LoginOutput, err error) {
	var hashedPassword string
	var id int

	err = r.db.QueryRowContext(ctx, "SELECT id, password from users WHERE phone = $1;", input.Phone).Scan(&id, &hashedPassword)
	if err != nil {
		log.Error("[Repository:Login] Failed to get user data, error"+err.Error(), " | context_id: ", ctx)
		return
	}

	if id == 0 || hashedPassword == "" {
		log.Error("[Repository:Login] User not found", " | context_id: ", ctx)
		err = errors.New("user not found")
		return
	}

	err = r.comparePasswords(ctx, hashedPassword, input.Password)
	if err != nil {
		return
	}

	output.ID = id
	return
}

func (r *UserRepository) hashAndSalt(ctx context.Context, pwd string) string {
	saltedPass := pwd + passwordSalt
	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPass), bcrypt.MinCost)
	if err != nil {
		log.Error("[Repository:hashAndSalt] error hashing password using bcrypt, error: "+err.Error(), " | context_id: ", ctx)
	}
	return string(hash)
}

func (r *UserRepository) comparePasswords(ctx context.Context, hashedPwd string, pwd string) (err error) {
	byteHash := []byte(hashedPwd)
	saltedPass := pwd + passwordSalt
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(saltedPass))
	if err != nil {
		log.Error("[Repository:hashAndSalt] error comparing bcrypt password, error: "+err.Error(), " | context_id: ", ctx)
	}

	return nil
}
