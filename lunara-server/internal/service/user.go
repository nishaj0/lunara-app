package service

import (
	"context"
	"errors"

	"github.com/nishaj0/lunara-app/lunara-server/internal/model"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/jwt"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"github.com/nishaj0/lunara-app/lunara-server/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx context.Context, req *model.RegisterRequest) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return nil, err
	}
	user := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashed),
		FullName:     req.FullName,
	}
	err = repository.CreateUser(ctx, user)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func LoginUser(ctx context.Context, req *model.LoginRequest) (string, *model.User, error) {
	user, err := repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}
	token, err := jwt.GenerateToken(user.ID, user.Email, req.RememberMe)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
