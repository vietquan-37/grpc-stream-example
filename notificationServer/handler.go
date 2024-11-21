package notificationServer

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vietquan-37/grpc-stream-example/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedNotificationServiceServer
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{
		redisClient: redisClient,
	}
}
func (h *Handler) GetNotification(req *pb.NotificationRequest, stream grpc.ServerStreamingServer[pb.NotificationResponse]) error {
	pubsub := h.redisClient.Subscribe(stream.Context(), fmt.Sprintf("notifications/%s", req.UserId))
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case msg := <-pubsub.Channel():
			if err := stream.SendMsg(&pb.NotificationResponse{
				UserId:   req.GetUserId(),
				Content:  fmt.Sprintf("New notification at : %s: %s", time.Now(), msg.Payload),
				CreateAt: timestamppb.New(time.Now()),
			}); err != nil {
				return fmt.Errorf("could not send notification: %w", err)
			}
		}
	}
}
