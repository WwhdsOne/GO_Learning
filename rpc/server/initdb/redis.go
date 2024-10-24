package initdb

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func Redis(password, address, port string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", address, port), // Redis 服务器地址
		Password: password,                            // Redis 密码，如果没有密码则留空
		DB:       0,                                   // 使用默认数据库
	})

	// 测试连接
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
