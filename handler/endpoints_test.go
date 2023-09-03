package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestServer_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := repository.NewMockUserInterface(ctrl)
	authRepo := repository.NewMockAuthInterface(ctrl)
	s := NewServer(NewServerOptions{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	})

	tests := map[string]struct {
		req        string
		mock       func()
		statusCode int
	}{
		"failed to bind request": {
			req: `{
				"name": "name",
				"password": "Password123!",
				"phone": 123
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"invalid password format": {
			req: `{
				"name": "name",
				"password": "password",
				"phone": "+62123123123"
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"failed when insert user to db": {
			req: `{
				"name": "name",
				"password": "Password123!",
				"phone": "+62123123123"
			  }`,
			mock: func() {
				userRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			statusCode: 400,
		},
		"success registration": {
			req: `{
				"name": "name",
				"password": "Password123!",
				"phone": "+62123123123"
			  }`,
			mock: func() {
				userRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)
			},
			statusCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.mock()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if err := s.Register(c); err != nil {
				t.Errorf("Server.Login() error = %v", err)
			}

			if tt.statusCode != rec.Code {
				t.Errorf("Server.Login() unexpected status code = %d, want = %d", rec.Code, tt.statusCode)
			}
		})
	}
}

func TestServer_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := repository.NewMockUserInterface(ctrl)
	authRepo := repository.NewMockAuthInterface(ctrl)
	s := NewServer(NewServerOptions{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	})

	tests := map[string]struct {
		req        string
		mock       func()
		statusCode int
	}{
		"failed to bind request": {
			req: `{
				"phone": 123,
				"password": "Password123!"
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"failed when insert user to db": {
			req: `{
				"phone": "+62123123123",
				"password": "Password123!"
			  }`,
			mock: func() {
				userRepo.EXPECT().Login(gomock.Any(), gomock.Any()).Return(repository.LoginOutput{}, errors.New("database error"))
			},
			statusCode: 400,
		},
		"failed to generate token": {
			req: `{
				"phone": "+62123123123",
				"password": "Password123!"
			  }`,
			mock: func() {
				userRepo.EXPECT().Login(gomock.Any(), gomock.Any()).Return(repository.LoginOutput{ID: 1}, nil)
				authRepo.EXPECT().GenerateToken(gomock.Any()).Return(repository.GenerateTokenOutput{}, errors.New("failed when parsing private key"))
			},
			statusCode: 400,
		},
		"success login": {
			req: `{
				"phone": "+62123123123",
				"password": "Password123!"
			  }`,
			mock: func() {
				userRepo.EXPECT().Login(gomock.Any(), gomock.Any()).Return(repository.LoginOutput{ID: 1}, nil)
				authRepo.EXPECT().GenerateToken(gomock.Any()).Return(repository.GenerateTokenOutput{Token: "token"}, nil)
			},
			statusCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.mock()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if err := s.Login(c); err != nil {
				t.Errorf("Server.Login() error = %v", err)
			}

			if tt.statusCode != rec.Code {
				t.Errorf("Server.Login() unexpected status code = %d, want = %d", rec.Code, tt.statusCode)
			}
		})
	}
}

func TestServer_GetUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := repository.NewMockUserInterface(ctrl)
	authRepo := repository.NewMockAuthInterface(ctrl)
	s := NewServer(NewServerOptions{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	})

	tests := map[string]struct {
		mock       func()
		statusCode int
	}{
		"failed when get user data from db": {
			mock: func() {
				userRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(repository.GetUserOutput{}, errors.New("database error"))
			},
			statusCode: 400,
		},
		"success get user data": {
			mock: func() {
				userRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(repository.GetUserOutput{Name: "name", Phone: "+628123123"}, nil)
			},
			statusCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.mock()

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/user", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set(CtxUserID, 1)

			if err := s.GetUserData(c); err != nil {
				t.Errorf("Server.GetUserData() error = %v", err)
			}

			if tt.statusCode != rec.Code {
				t.Errorf("Server.GetUserData() unexpected status code = %d, want = %d", rec.Code, tt.statusCode)
			}
		})
	}
}

func TestServer_UpdateUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := repository.NewMockUserInterface(ctrl)
	authRepo := repository.NewMockAuthInterface(ctrl)
	s := NewServer(NewServerOptions{
		UserRepository: userRepo,
		AuthRepository: authRepo,
	})

	tests := map[string]struct {
		req        string
		mock       func()
		statusCode int
	}{
		"failed to bind request": {
			req: `{
				"name": "name",
				"phone": 123
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"missing field for update data": {
			req: `{}`,
			mock: func() {
			},
			statusCode: 400,
		},
		"invalid name format": {
			req: `{
				"name": "a",
				"phone": "+62123123123"
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"invalid phone format": {
			req: `{
				"name": "name",
				"phone": "123"
			  }`,
			mock: func() {
			},
			statusCode: 400,
		},
		"failed when update user data in db": {
			req: `{
				"name": "name",
				"phone": "+62123123123"
			  }`,
			mock: func() {
				userRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			statusCode: 400,
		},
		"success update user data": {
			req: `{
				"name": "name",
				"phone": "+62123123123"
			  }`,
			mock: func() {
				userRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			statusCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.mock()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(tt.req))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set(CtxUserID, 1)

			if err := s.UpdateUserData(c); err != nil {
				t.Errorf("Server.GetUserData() error = %v", err)
			}

			if tt.statusCode != rec.Code {
				t.Errorf("Server.GetUserData() unexpected status code = %d, want = %d", rec.Code, tt.statusCode)
			}
		})
	}
}
