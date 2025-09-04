package user

import (
	"context"

	"github.com/Rocksus/fundtract/internal/model/dto"
)

func (uu *userUsercase) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	return dto.LoginResponse{}, nil
}
