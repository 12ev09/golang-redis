package repository

import (
	"context"
	"encoding/json"

	"github.com/12ev09/golang-redis/models"
	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

// Redisに接続する
func SetupRedis() {
	Cache = redis.NewClient(
		&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		},
	)
}

// redisからユーザーリストを取得
func GetUserList(uuid string) (models.Users, error) {
	// キャッシュから指定したuuidで値を取得
	data, err := Cache.Get(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}

	userList := new(models.Users)
	// json形式でとってきたデータをUser型に変換
	if err := json.Unmarshal([]byte(data), userList); err != nil {
		return nil, err
	}

	return *userList, nil
}
