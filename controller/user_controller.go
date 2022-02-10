package controller

import (
	"net/http"

	"github.com/12ev09/golang-redis/controller_interface"
	"github.com/12ev09/golang-redis/repository_interface"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepository repository_interface.UserRepositoryInterface
}

func NewUserController(userRepository repository_interface.UserRepositoryInterface) controller_interface.UserControllerInterface {
	return UserController{
		UserRepository: userRepository,
	}
}

func (self UserController) GetUserList(c *gin.Context) {
	// リクエストからuuidを取得
	uuid := c.Param("uuid")

	//redisからデータを取得
	userList, err := self.UserRepository.GetUserList(uuid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, userList)
}
