package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"internal/model"
	"time"
)

type RedisDB struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisDB(address, password string, db int, ttl time.Duration) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return &RedisDB{client: client, ttl: ttl}, nil
}

func (r *RedisDB) StoreIdempotencyToken(generatedKey string, rd *model.ResponseData) error {
	jsonData, err := json.Marshal(rd)
	if err != nil {
		return err
	}

	err = r.client.Set(generatedKey, jsonData, r.ttl).Err()
	if err != nil {
		return err
	}

	err = r.client.Expire(generatedKey, r.ttl).Err()
	return err
}

func (r *RedisDB) RetrieveIdempotencyToken(generatedKey string) (*model.ResponseData, error) {
	jsonData, err := r.client.Get(generatedKey).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var response model.ResponseData
	err = json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return nil, err // Error occurred while unmarshalling the data
	}

	return &response, nil
}

func (r *RedisDB) GenerateTokenKey(idempotencyToken string, expression string, userId int64) string {
	return fmt.Sprintf("%s__%s__%d", idempotencyToken, expression, userId)
}
