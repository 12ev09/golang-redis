# go-redis

[go-redis](https://github.com/go-redis/redis)を使用してみた

## ファイル構成
```
.
├── controller
│   └── user_controller.go
├── controller_interface
│   └── user_controller_interface.go
├── repository
│   ├── redis.go
│   └── user_repository.go
├── repository_interface
│   └── user_repository_interface.go
├── models
│   └── user.go
├── command
│   └── testdata.go
├── main.go
├── go.mod
├── go.sum
├── docker-compose.yml
├── Dockerfile
├── MakeFile
└── README.md
```

## 動かし方

コンテナをbuild
```
make build
```

コンテナを起動
```
make up
```

ダミーデータをredisに投入
```
make testdata
```

uuidからユーザーを取得
```
http:localhost:8080/users/[uuid]
```

コンテナを停止
```
make down
```

## 参考にしたサイト

- https://selfnote.work/20210719/programming/golang/golang-redis/
- https://github.com/go-redis/redis