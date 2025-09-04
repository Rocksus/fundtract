package user

import (
	"context"

	"github.com/Rocksus/fundtract/internal/model"
	"github.com/Rocksus/fundtract/internal/model/dto"
)

type Usecase interface {
	GetUserById(ctx context.Context, userId int) (model.UserData, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

type userUsercase struct {
}

func NewUsecase() Usecase {
	return &userUsercase{}
}
