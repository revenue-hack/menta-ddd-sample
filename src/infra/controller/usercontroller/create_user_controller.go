package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revenue-hack/menta-ddd-sample/src/app/userapp"
	"github.com/revenue-hack/menta-ddd-sample/src/infra/rdb/repoimpl"
)

type createUserController struct {
}

func newCreateUserController() *createUserController {
	return &createUserController{}
}

func (c *createUserController) exec(ctx *gin.Context) {
	var in userapp.CreateUserRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	userRepoImpl := repoimpl.NewUserRepositoryImpl()

	if err := userapp.NewCreateUserAppService(userRepoImpl).Exec(ctx, &in); err != nil {
		ctx.Error(err)
		return
	}
}
