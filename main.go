package main

import (
	"net/http"

	"github.com/12ev09/golang-redis/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	//redis
	repository.SetupRedis()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.GET("users/:uuid", getUserList)

	router.Run(":8080")
}

func getUserList(c *gin.Context) {
	// リクエストからuuidを取得
	uuid := c.Param("uuid")

	//redisからデータを取得
	userList, err := repository.GetUserList(uuid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, userList)
}
