/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/2/8 3:37 下午
 */

package redis

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
)

var client *redis.Client

// InitRedis
func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
}

// TestGO_Redis
func TestGO_Redis(t *testing.T) {
	InitRedis()
	bytes, err := proto.Marshal(&Person{
		Name: "123",
	})
	if err != nil {
		fmt.Println("err ", err.Error())
		panic(err)
	}
	fmt.Println("bb: ", bytes)
	HashSet("z", "w", bytes)
	HashGet("z", "w")
	HashGetAll("z")
}

func HashSet(key, field string, val []byte) error {
	cmd := client.HSet(key, field, val)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

func HashGet(key, field string) error {
	cmd := client.HGet(key, field)
	r, err := cmd.Result()
	if err != nil {
		return err
	}
	fmt.Println("result :", strings.TrimLeft(r, "\n"))
	return nil
}

// HashGetAll
func HashGetAll(key string) error {
	cmd := client.HGetAll(key)
	m, err := cmd.Result()
	if err != nil {
		return err
	}
	for i, v := range m {
		str := fmt.Sprintf("key : %v, value : %v", i, strings.TrimLeft(v,"\n"))
		fmt.Println(str)
		fmt.Println(strings.TrimLeft(strings.TrimLeft(v, "\n"), " ") == "123")
	}
	return nil
}
