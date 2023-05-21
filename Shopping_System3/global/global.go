package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// 定义全局变量
var (
	Xdb         *sqlx.DB
	RedisClient *redis.Client
)
