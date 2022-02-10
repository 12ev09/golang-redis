package controller_interface

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	GetUserList(c *gin.Context)
}
