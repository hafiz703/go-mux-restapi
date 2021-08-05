package cache

import (
	"encoding/json"
	"errors"
	"time"

	"golang-mux-api/entity"

	"fmt"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) UserCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, user *entity.User) {
	client := cache.getClient()

	// serialize user object to JSON
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.User {
	client := cache.getClient()
	 
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	user := entity.User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}

	return &user
}

func (cache *redisCache) GetAll() ([]entity.User) {
	fmt.Println("GETTING ALL")
	client := cache.getClient()
	
	// rows := client.Scan(0, "keys *", 0).Iterator()
	rows,_ := client.Keys("*").Result()
 
 
	fmt.Println("rows:",rows)
	var users []entity.User = []entity.User{}
	for _, key := range rows {

		user := *cache.Get(key)
		users = append(users, user)
    }
 
	 
	return users

	 
}

func (cache *redisCache) Validate(user *entity.User) error {
	if user == nil {
		err := errors.New("The user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("The user name is empty")
		return err
	}
	return nil
}

func (cache *redisCache) Delete(key string) {
	client := cache.getClient()
	 
	client.Del(key).Result()
	 
}
 
