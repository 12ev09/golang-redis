package repository_interface

import "github.com/12ev09/golang-redis/models"

type UserRepositoryInterface interface {
	GetUserList(id string) (*models.Users, error)
}
