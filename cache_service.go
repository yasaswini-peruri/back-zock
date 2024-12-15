package services

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"productmanagement/models"
	"context"
)

var rdb *redis.Client

func InitCache() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func CacheProduct(id string, product *models.Product) {
	data, _ := json.Marshal(product)
	rdb.Set(context.Background(), "product:"+id, data, 0)
}

func GetCachedProduct(id string) (*models.Product, error) {
	data, err := rdb.Get(context.Background(), "product:"+id).Result()
	if err != nil {
		return nil, err
	}
	product := &models.Product{}
	json.Unmarshal([]byte(data), product)
	return product, nil
}
