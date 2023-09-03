package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
)

const (
	CtxUserID = "id"
)

func CreateMiddleware(v repository.AuthInterface) ([]echo.MiddlewareFunc, error) {
	spec, err := generated.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	validator := middleware.OapiRequestValidatorWithOptions(spec,
		&middleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}

func NewAuthenticator(v repository.AuthInterface) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(v, ctx, input)
	}
}

func Authenticate(v repository.AuthInterface, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "BearerAuth" {
		return fmt.Errorf("security scheme %s != 'BearerAuth'", input.SecuritySchemeName)
	}

	req := input.RequestValidationInput.Request

	authHeader := req.Header.Get("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return errors.New("Invalid Auth")
	}
	authToken := strings.TrimPrefix(authHeader, prefix)
	output, err := v.ParseToken(repository.ParseTokenInput{
		Token: authToken,
	})
	if err != nil {
		return errors.New("Invalid Auth")
	}

	expTime := time.Unix(output.Exp, 0)
	if time.Now().After(expTime) {
		return errors.New("Token Expired")
	}

	eCtx := middleware.GetEchoContext(ctx)
	eCtx.Set(CtxUserID, output.ID)

	return nil
}
