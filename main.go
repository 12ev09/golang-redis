package main

import (
	"github.com/12ev09/golang-redis/controller"
	"github.com/12ev09/golang-redis/controller_interface"
	"github.com/12ev09/golang-redis/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	//redis
	repository.SetupRedis()

	router := gin.Default()

	var userController controller_interface.UserControllerInterface = controller.NewUserController(
		repository.NewUserRepository(),
	)

	router.GET("users/:uuid", userController.GetUserList)

	router.Run(":8080")
}
