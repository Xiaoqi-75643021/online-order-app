package service

import (
	"online-ordering-app/internal/model"
	"online-ordering-app/internal/repository"
	"online-ordering-app/internal/utils"
	"online-ordering-app/pkg/jwt"
)

func Login(username, password string) (string, error) {
	authenticated, err := repository.AuthenticateUser(username, password)
	if err != nil {
		return "", err
	}
	if !authenticated {
		return "", ErrInvalidCredentails
	}

	user, err := repository.FindUserByName(username)
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Register(username, password, role string) (string, error) {
	_, err := repository.FindUserByName(username)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", ErrHashPasswordFailed
	}

	user := &model.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}
	err = repository.CreateUser(user)
	if err != nil {
		return "", ErrCreateUserFailed
	}
	token, err := jwt.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return "", ErrGenerateTokenFailed
	}
	return token, nil
}
