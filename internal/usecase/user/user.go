package user

import (
	"context"

	"github.com/Rocksus/fundtract/internal/model"
)

func (uu *userUsercase) GetUserById(ctx context.Context, userId int) (model.UserData, error) {
	return model.UserData{}, nil
}
