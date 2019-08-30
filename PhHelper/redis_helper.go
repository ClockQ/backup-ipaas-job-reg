package PhHelper

import "github.com/go-redis/redis"

type PhRedisHelper struct {
	Redis    *redis.Client
	host     string
	port     string
	password string
}

func (helper *PhRedisHelper) New(host, port, pwd string) *PhRedisHelper {
	helper.host = host
	helper.port = port
	helper.password = pwd

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: helper.password,
	})
	helper.Redis = client

	return helper
}
