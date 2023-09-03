package main

import (
	"os"
	"strconv"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbDsn := os.Getenv("DATABASE_URL")
	var userRepo repository.UserInterface = repository.NewUserRepository(repository.NewUserRepositoryOptions{
		Dsn: dbDsn,
	})

	privateKey := os.Getenv("JWT_PRIVATE_KEY")
	publicKey := os.Getenv("JWT_PUBLIC_KEY")
	tokenExpiry := os.Getenv("TOKEN_EXPIRY") // minutes

	exp, err := strconv.Atoi(tokenExpiry)
	if err != nil {
		panic(err)
	}

	var authRepo repository.AuthInterface = repository.NewJWTRepository(repository.NewJWTRepositoryOptions{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		ExpiryTime: time.Minute * time.Duration(exp),
	})

	middleware, err := newMiddleware(authRepo)
	if err != nil {
		panic(err)
	}
	e.Use(middleware...)

	var server generated.ServerInterface = newServer(userRepo, authRepo)

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer(userRepo repository.UserInterface, authRepo repository.AuthInterface) *handler.Server {
	opts := handler.NewServerOptions{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
	return handler.NewServer(opts)
}

func newMiddleware(authRepo repository.AuthInterface) ([]echo.MiddlewareFunc, error) {
	return handler.CreateMiddleware(authRepo)
}
