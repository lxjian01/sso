package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sso/config"
	"sso/log"
	"time"
)

var (
	client *redis.Client
	ctx = context.Background()
)

// 初始化redis数据库
func InitRedisPool(c *config.RedisConfig) {
	conStr := fmt.Sprintf("%s,%d",c.Host, c.Port)
	ops := &redis.Options{
		Addr:     conStr,
		DB:       0,  // use default DB
		PoolSize: c.PoolSize,
		MaxRetries: c.MaxRetries,
		DialTimeout: (60 * time.Second),
		IdleTimeout: (60 * time.Second),
	}
	client = redis.NewClient(ops)
}

func Set(key string, value interface{}, expiration time.Duration){
	client.Set(ctx,key,value,expiration)
}

// 向key的hash中添加元素field的值
func HashSet(key string, field string, data interface{}) {
	err := client.HSet(ctx,key, field, data)
	if err != nil {
		log.Error("Redis HSet Error:", err)
	}
}

func SetExpire(key string,expiration time.Duration) bool {
	val,err := client.Expire(ctx,key,expiration).Result()
	if err != nil {
		log.Error("Redis Expire Error:", err)
	}

	return val
}

// 批量向key的hash添加对应元素field的值
func BatchHashSet(key string, fields map[string]interface{}) bool {
	val, err := client.HMSet(ctx,key, fields).Result()
	if err != nil {
		log.Error("Redis HMSet Error:", err)
	}

	return val
}

// 通过key获取hash的元素值
func HashGet(key string, field string) string {
	result := ""
	val, err := client.HGet(ctx,key, field).Result()
	if err == redis.Nil {
		log.Error("Key Doesn't Exists:", field)
		return result
	}else if err != nil {
		log.Error("Redis HGet Error:", err)
		return result
	}
	return val
}

func CloseRedis(){
	if err := client.Close();err != nil{
		log.Error("Close redis error by",err)
	}
}
