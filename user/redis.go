package user

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type SimpleRedis interface {
	// Close()
	Get(key string) *User
	Put(key string, value *User) error
}

type simpleRedis struct {
	ctx    context.Context
	client *redis.Client
	close  func()
}

func NewSimpleRedis() SimpleRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "", // no password set
		DB:       0,  // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})
	return &simpleRedis{
		ctx:    context.Background(),
		client: client,
		// close:  func() { client.Close() },
	}
}

func (simpleRedis *simpleRedis) Get(key string) *User {
	value, err := simpleRedis.client.Get(simpleRedis.ctx, key).Result()
	if err == redis.Nil {
		log.Printf("%s does not exist", key)
	} else if err != nil {
		log.Printf("error=%v", err)
	} else {
		log.Printf("redis found key=%s", key)
		var user User
		json.Unmarshal([]byte(value), &user)
		return &user
	}
	return nil
}

func (simpleRedis *simpleRedis) Put(key string, user *User) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = simpleRedis.client.Set(simpleRedis.ctx, key, string(value), 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return err
}
