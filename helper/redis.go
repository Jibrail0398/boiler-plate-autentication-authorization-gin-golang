package helper

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
	});

	return client
}



func StoreWithTime(key string, value string, timelong time.Duration, client *redis.Client) error {
	
	ctx :=context.Background()
	
	timelong = time.Duration(timelong * time.Second)

	err := client.Set(ctx, key, value, timelong).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetDataRedis(key string, client *redis.Client) (string,error){
	ctx := context.Background()
	redisGet := client.Get(ctx,key)

	if err:=redisGet.Err();err!=nil{
		
		return "nil",err
	}

	res, err := redisGet.Result()
    if err != nil {
        
        return "",err
    }

	return res,nil

	
}