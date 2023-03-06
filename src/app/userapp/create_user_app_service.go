package userapp

import (
	"context"

	"github.com/revenue-hack/menta-ddd-sample/src/domain/userdm"
)

type createUserAppService struct {
	userRepo userdm.UserRepository
}

type CreateUserRequest struct {
	FirstName string
	LastName  string
}

func NewCreateUserAppService(userRepo userdm.UserRepository) *createUserAppService {
	return &createUserAppService{userRepo}
}

func (app *createUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	user, err := userdm.GenWhenCreate(req.FirstName, req.LastName)
	if err != nil {
		return err
	}

	return app.userRepo.Store(ctx, user)
}
