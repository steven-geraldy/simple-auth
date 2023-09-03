package handler

import "github.com/SawitProRecruitment/UserService/repository"

type Server struct {
	UserRepository repository.UserInterface
	AuthRepository repository.AuthInterface
}

type NewServerOptions struct {
	UserRepository repository.UserInterface
	AuthRepository repository.AuthInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserRepository: opts.UserRepository,
		AuthRepository: opts.AuthRepository,
	}
}
