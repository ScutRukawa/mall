package utils

import (
	"context"
	"errors"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient() *RedisClient {
	redisIns := &RedisClient{}
	redisIns.rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return redisIns
}
func (r *RedisClient) Aquire(ctx context.Context, key string, lockId string) error {
	_, err := r.rdb.SetNX(ctx, key, lockId, 10*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Release(ctx context.Context, key string, lockId string) error {
	// v, err := r.rdb.Get(ctx, key).Result()
	// if err != nil {
	// 	return err
	// }
	// if v == lockId {
	// 	r.rdb.Del(ctx, key)
	// 	return nil
	// } else {
	// 	return errors.New("当前没持有锁，不能删除不属于自己的锁")
	// }
	script :=
		"if (redis.call('get',KEYS[1])==ARGV[1]) then \n" +
			"redis.call('del',KEYS[1]) \n" +
			"return 1\n" +
			"else\n" +
			"return 0\n" +
			"end\n"
	val, _ := r.rdb.Eval(ctx, script, []string{key}, lockId).Result()
	if v, _ := val.(int64); v == 1 {
		return nil
	}
	return errors.New("当前没持有锁，不能删除不属于自己的锁")
}
