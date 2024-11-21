package main

import (
	"context"
	"io"
	"log"

	"github.com/vietquan-37/grpc-stream-example/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:5051", opts...)
	if err != nil {
		log.Fatalf("cannot connect to grpc: %v", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)
	stream, err := client.GetNotification(context.Background(), &pb.NotificationRequest{
		UserId: "123",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		notification, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNotification(_) = _, %v", client, err)
		}
		log.Println(notification)
	}
}
