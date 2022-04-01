package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	return client
}

func getDataRedis(con *redis.Client, key string) {
	fmt.Println("Get Data")
	val, err := con.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Data Dengan Key '", key, "' Kosong")
		return
	}

	fmt.Println("GET data untuk key '", key, "' : ", val)
}

func TestGoRedis() {

	var redisHost = "localhost:6739"
	var redisPassword = ""

	con := newRedisClient(redisHost, redisPassword)

	fmt.Println("Inisialisasi Redis")

	key := "pbp"
	data := 99
	exp := time.Duration(3) * time.Second

	ctx := context.Background()

	fmt.Println("Set Nilai PBP")
	err := con.Set(ctx, key, data, exp).Err()
	if err != nil {
		fmt.Printf("Gagal SET data. error: %v", err)
		return
	}

	fmt.Println("SET data untuk key '", key, "' Berhasil.")

	getDataRedis(con, key)

	fmt.Println("Mencoba Increment Nilai PBP")
	err = con.Incr(ctx, key).Err()
	if err != nil {
		fmt.Printf("Gagal Increment data. error: %v", err)
		return
	}
	fmt.Println("Increment data untuk key '", key, "' Berhasil.")

	getDataRedis(con, key)

	err = con.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Gagal Delete key. error: %v", err)
		return
	}
	fmt.Println("Delete key '", key, "' Berhasil.")

	getDataRedis(con, key)
}
