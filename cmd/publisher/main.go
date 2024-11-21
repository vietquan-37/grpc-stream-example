package main

import (
	"context"
	"fmt"
	"log"
	"time"

	grpcstream "github.com/vietquan-37/grpc-stream-example"
)

func main() {
	redis := grpcstream.NewRedisClient(context.Background())
	channelName := fmt.Sprintf("notifications/%s", "123")
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-context.Background().Done():
			log.Fatal("shutting down")
		case t := <-ticker.C:
			if err := redis.Publish(context.Background(), channelName, fmt.Sprintf("New message %s", t.String())).Err(); err != nil {
				panic(err)
			}
		}
	}

}
