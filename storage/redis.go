package storage

import (
	"distributed_calculator/config"
	"distributed_calculator/model"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

type RedisDB struct {
	client *redis.Client
	ttl    time.Duration
}

func Redis(cfg *config.Config) (*RedisDB, error) {
	redisDB := &RedisDB{}
	err := redisDB.Init(cfg)

	if err != nil {
		return nil, err
	}

	return redisDB, nil
}

func (r *RedisDB) Init(cfg *config.Config) error {
	r.ttl = cfg.QuickAccessStorage.TTL
	r.client = redis.NewClient(&redis.Options{
		Addr:     cfg.QuickAccessStorage.Address,
		Password: cfg.QuickAccessStorage.Password,
		DB:       cfg.QuickAccessStorage.DB,
	})

	_, err := r.client.Ping().Result()

	return err
}

func (r *RedisDB) StoreIdempotencyToken(idempotencyToken string, expression string, responseData *model.ResponseData) error {
	token := generateToken(idempotencyToken, expression)

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return err
	}

	err = r.client.Set(token, jsonData, r.ttl).Err()
	if err != nil {
		return err
	}

	err = r.client.Expire(token, r.ttl).Err()
	return err
}

func (r *RedisDB) RetrieveIdempotencyToken(idempotencyToken string, expression string) (*model.ResponseData, error) {
	token := generateToken(idempotencyToken, expression)

	jsonData, err := r.client.Get(token).Result()
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

func generateToken(idempotencyToken string, expression string) string {
	return idempotencyToken + "__" + expression
}
