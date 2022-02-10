package repository

import (
	"context"
	"encoding/json"

	"github.com/12ev09/golang-redis/models"
	"github.com/12ev09/golang-redis/repository_interface"
)

type UserRepository struct{}

func NewUserRepository() repository_interface.UserRepositoryInterface {
	return UserRepository{}
}

func (self UserRepository) GetUserList(id string) (*models.Users, error) {
	// キャッシュから指定したuuidで値を取得
	data, err := Cache.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}

	userList := new(models.Users)
	// json形式でとってきたデータをUser型に変換
	if err := json.Unmarshal([]byte(data), userList); err != nil {
		return nil, err
	}

	return userList, nil
}
