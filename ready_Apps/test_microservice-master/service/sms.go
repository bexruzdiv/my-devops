package service

import (
	"context"
	pb "genproto/sms_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/sms_service/storage"
)

// SendService struct...
type SendService struct {
	storage storage.StorageI
}

// NewSendService ...
func NewSendService(db *sqlx.DB) *SendService {
	return &SendService{storage: storage.NewStoragePg(db)}
}

//Send ...
func (s *SendService) Send(ctx context.Context, req *pb.Sms) (*gpb.Empty, error) {
	err := s.storage.SendSms().Send(req.ShipperId, req.Text, req.Recipients)
	if err != nil {
		return &gpb.Empty{}, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}
