package usercontroller

import "github.com/gin-gonic/gin"

func NewUserController(g *gin.Engine) {
	g.POST("/users", func(ctx *gin.Context) {
		newCreateUserController().exec(ctx)
	})
}
