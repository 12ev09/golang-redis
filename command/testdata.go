package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/12ev09/golang-redis/models"
	"github.com/bxcodec/faker/v3"
	"github.com/go-redis/redis/v8"
)

type SomeStructWithTags struct {
	Email     string `faker:"email"`
	Name      string `faker:"name"`
	UUID      string `faker:"uuid_digit"`
	AccountID int    `faker:"oneof:15, 27, 61"`
}

func SetupRedis() *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		},
	)
}

// redisにデータを入れて、取得するためのuuidを発行する
func main() {
	a := SomeStructWithTags{}
	err := faker.FakeData(&a)
	if err != nil {
		panic(err)
	}

	var userList models.Users
	for i := 0; i < 20; i++ {
		if err := faker.FakeData(&a); err != nil {
			panic(err)
		}
		userList = append(userList, models.User{
			AccountID: a.AccountID,
			Name:      a.Name,
			Email:     a.Email,
		})
	}

	// Redisに格納するために、シリアライズ
	serialize, err := json.Marshal(&userList)
	if err != nil {
		panic(err)
	}

	// UUIDはアクセスするために必要
	UUID := a.UUID
	fmt.Println("UUID: ", UUID)

	// Redisに接続
	c := context.Background()
	r := SetupRedis()

	if err := r.Set(c, UUID, serialize, time.Hour*24).Err(); err != nil {
		panic(err)
	}

	log.Println("complete")
}
