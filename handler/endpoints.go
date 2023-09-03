package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (s *Server) Register(c echo.Context) error {
	var req generated.RegisterJSONRequestBody
	var resp generated.Response

	ctx := context.Background()

	err := c.Bind(&req)
	if err != nil {
		log.Error("[API:Register] failed to bind request", " | context_id: ", ctx)
		resp.Message = "Failed to bind Request"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !validateUserName(req.Name) || !validateUserPassword(req.Password) || !validateUserPhone(req.Phone) {
		log.Error("[API:Register] Invalid Request Format", " | context_id: ", ctx)
		resp.Message = "Invalid Request Format"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = s.UserRepository.Register(ctx, repository.RegisterUserInput{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		log.Error("[API:Register] failed to Register user, error:", err.Error(), " | context_id: ", ctx)
		resp.Message = "Registration Failed"
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp.Message = "Registration Success"
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) Login(c echo.Context) error {
	var req generated.LoginJSONRequestBody
	var resp generated.LoginResponse

	ctx := context.Background()

	err := c.Bind(&req)
	if err != nil {
		log.Error("[API:Login] failed to bind request", " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "Login Failed",
		})
	}

	loginOutput, err := s.UserRepository.Login(ctx, repository.LoginInput{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		log.Error("[API:Login] failed to Register user, error: ", err.Error(), " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "Login Failed",
		})
	}

	output, err := s.AuthRepository.GenerateToken(repository.GenerateTokenInput{
		ID: loginOutput.ID,
	})
	if err != nil {
		log.Error("[API:Login] failed to generate access token, error: ", err.Error(), " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "Login Failed",
		})
	}

	resp.Message = fmt.Sprintf("Login Success")
	resp.Token = output.Token
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) GetUserData(c echo.Context) error {
	var resp generated.UserResponse

	ctx := context.Background()
	id := c.Get(CtxUserID).(int)

	user, err := s.UserRepository.GetUser(ctx, repository.GetUserInput{
		ID: id,
	})
	if err != nil {
		log.Error("[API:GetUserData] failed to get user, error:", err.Error(), " | context_id: : ", ctx)
		return c.JSON(http.StatusBadRequest, generated.UserResponse{
			Message: "Failed Get User Data",
		})
	}

	resp.Message = fmt.Sprintf("Success Get User Data")
	resp.Name = user.Name
	resp.Phone = user.Phone
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateUserData(c echo.Context) error {
	var req generated.UpdateUserDataJSONRequestBody
	var resp generated.Response

	ctx := context.Background()
	id := c.Get(CtxUserID).(int)

	err := c.Bind(&req)
	if err != nil {
		log.Error("[API:UpdateUserData] failed to bind request", " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "UpdateUserData Failed",
		})
	}

	if req.Phone == nil && req.Name == nil {
		log.Error("[API:UpdateUserData] no data to update", " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "UpdateUserData Failed",
		})
	}

	input := repository.UpdateUserInput{
		ID: id,
	}
	if req.Name != nil {
		if !validateUserName(*req.Name) {
			log.Error("[API:UpdateUserData] invalid name format", " | context_id: ", ctx)
			return c.JSON(http.StatusBadRequest, generated.Response{
				Message: "UpdateUserData Failed",
			})
		}
		input.Name = *req.Name
	}
	if req.Phone != nil {
		if !validateUserPhone(*req.Phone) {
			log.Error("[API:UpdateUserData] invalid phone format", " | context_id: ", ctx)
			return c.JSON(http.StatusBadRequest, generated.Response{
				Message: "UpdateUserData Failed",
			})
		}
		input.Phone = *req.Phone
	}

	err = s.UserRepository.UpdateUser(ctx, input)
	if err != nil {
		log.Error("[API:UpdateUserData] failed to Update user, error:", err.Error(), " | context_id: ", ctx)
		return c.JSON(http.StatusBadRequest, generated.Response{
			Message: "UpdateUserData Failed",
		})
	}

	resp.Message = fmt.Sprintf("Update User Success")
	return c.JSON(http.StatusOK, resp)
}
