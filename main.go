package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func acquireLock(client *redis.Client, lockKey string, timeout time.Duration) bool {
	ctx := context.Background()

	// Try to acquire the lock with SETNX command (SET if Not eXist)
	lockAcquired, err := client.SetNX(ctx, lockKey, "1", timeout).Result()
	if err != nil {
		fmt.Println("Error acquiring lock")
		return false
	}

	return lockAcquired

}

func releaseLock(client *redis.Client, lockKey string) {
	ctx := context.Background()
	client.Del(ctx, lockKey)
}

func main() {
	// Create a redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer client.Close()

	// Define the lockKey and lockTimeout
	lockKey := "my_lock"
	lockTimeout := 20 * time.Second

	if acquireLock(client, lockKey, lockTimeout) {
		fmt.Println("Lock acquired successfully!")
		// simulate some work with the lock
		time.Sleep(20 * time.Second)
		fmt.Println("Work Done!")

		// release the lock to free resource
		releaseLock(client, lockKey)
		fmt.Println("Lock released!")
	} else {
		fmt.Println("Failed to acquire lock. The resource is already locked!")
	}

}
