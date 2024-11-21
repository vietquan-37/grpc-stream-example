package main

import (
	"context"
	"log"
	"net"

	grpcstream "github.com/vietquan-37/grpc-stream-example"
	"github.com/vietquan-37/grpc-stream-example/notificationServer"
	"github.com/vietquan-37/grpc-stream-example/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalf("error while listening to port 5051: %v", err)
	}
	redisClient := grpcstream.NewRedisClient(context.Background())
	handler := notificationServer.NewHandler(redisClient)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNotificationServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve : %v", err)
	}
}
