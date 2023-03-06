package controller

import (
	"context"

	"github.com/revenue-hack/menta-ddd-sample/src/app/userapp"
	"github.com/revenue-hack/menta-ddd-sample/src/domain/userdm"
)

type CreateUserController struct {
	userRepo userdm.UserRepository
}

func NewCreateUserController(userRepo userdm.UserRepository) *CreateUserController {
	return &CreateUserController{
		userRepo: userRepo,
	}
}

func (c *CreateUserController) Exec(ctx context.Context, req *userapp.CreateUserRequest) error {
	app := userapp.NewCreateUserAppService(c.userRepo)
	return app.Exec(ctx, &userapp.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
}

/*
func (c *userController) GetUserByID(ctx context.Context, in *userinput.GetUserByIDInput) error {
	usecase := userusecase.NewGetUserByID(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.UserByID(out)
	return nil
}

func (c *userController) CreateUser(ctx context.Context, in *userinput.CreateUserInput) error {
	usecase := userusecase.NewCreateUser(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Create(out)
	return nil
}
*/
