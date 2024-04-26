package service

import (
	"context"
	"database/sql"
	pb "genproto/sms_service"

	gpb "github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	l "bitbucket.org/alien_soft/sms_service/pkg/logger"
	"bitbucket.org/alien_soft/sms_service/storage"
)

type SmsProviderService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewSmsProviderService(db *sqlx.DB, log l.Logger) *SmsProviderService {
	return &SmsProviderService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *SmsProviderService) Get(ctx context.Context, req *pb.ShipperId) (*pb.SmsProvider, error) {
	smsProvider, err := s.storage.SmsProvider().Get(req.ShipperId)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting sms provider, Not found", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting sms provider", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return smsProvider, nil
}

func (s *SmsProviderService) Create(ctx context.Context, req *pb.SmsProvider) (*gpb.Empty, error) {
	err := s.storage.SmsProvider().Create(req)
	if err != nil {
		s.logger.Error("Error while creating sms provider", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &gpb.Empty{}, nil
}

func (s *SmsProviderService) Delete(ctx context.Context, req *pb.ShipperId) (*gpb.Empty, error) {
	err := s.storage.SmsProvider().Delete(req.ShipperId)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while deleting sms provider, Not found", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while deleting sms provider", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &gpb.Empty{}, nil
}
