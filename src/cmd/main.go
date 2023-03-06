package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/revenue-hack/menta-ddd-sample/src/infra/controller/usercontroller"
	"github.com/revenue-hack/menta-ddd-sample/src/infra/middleware"
)

func main() {
	g := gin.Default()

	g.Use(middleware.HandleErrorMiddleware())

	usercontroller.NewUserController(g)

	log.Fatal(g.Run(":8000"))
}
