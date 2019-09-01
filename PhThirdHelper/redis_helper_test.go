package PhThirdHelper

import (
	"github.com/hashicorp/go-uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPhRedisHelper(t *testing.T) {

	const(
		RedisHost = ""
		RedisPort = ""
		RedisPwd  = ""
	)

	rh := PhRedisHelper{}.New(RedisHost, RedisPort, RedisPwd)
	id, _ := uuid.GenerateUUID()

	Convey("测试 Redis 写入 Hash", t, func() {
		err := rh.Redis.HSet(id, "test", 0).Err()
		So(err, ShouldBeNil)
	})

	Convey("测试 Redis 读取 Hash", t, func() {
		stepStr, err := rh.Redis.HGet(id, "test").Result()
		So(err, ShouldBeNil)
		So(stepStr, ShouldEqual, "0")
	})

	Convey("测试 Redis 删除", t, func() {
		err := rh.Redis.Del(id).Err()
		So(err, ShouldBeNil)
	})
}
