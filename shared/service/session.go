package service

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// Function Set Session
func SetSession(userID string, token string) {
	rdb.Set(context.Background(), userID, token, 24*time.Hour)
}

// Function Get Session
func GetSession(userID string) string {
	val, err := rdb.Get(context.Background(), userID).Result()
	if err != nil {
		return ""
	}
	return val
}
