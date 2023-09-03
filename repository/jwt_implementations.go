package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _ AuthInterface = &JWTRepository{}

func (r *JWTRepository) GenerateToken(input GenerateTokenInput) (output GenerateTokenOutput, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(r.privateKey))
	if err != nil {
		return
	}

	t := jwt.New(jwt.SigningMethodRS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = strconv.Itoa(input.ID)
	claims["exp"] = strconv.FormatInt(time.Now().Add(r.expiryTime).Unix(), 10)

	output.Token, err = t.SignedString(key)
	return
}

func (r *JWTRepository) ParseToken(input ParseTokenInput) (output ParseTokenOutput, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(r.publicKey))
	if err != nil {
		return
	}

	parsedToken, err := jwt.Parse(input.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Invalid Token")
		}
		return key, nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("Invalid Token")
		return
	}

	id, ok := claims["id"].(string)
	if !ok {
		err = errors.New("Invalid Token")
		return
	}

	output.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("Invalid Token")
		return
	}

	exp, ok := claims["exp"].(string)
	if !ok {
		err = errors.New("Invalid Token")
		return
	}

	output.Exp, err = strconv.ParseInt(exp, 10, 64)
	if err != nil {
		err = errors.New("Invalid Token")
		return
	}

	return
}
