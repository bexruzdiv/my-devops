package service

import (
	"context"
	"database/sql"
	pb "genproto/sms_service"

	l "bitbucket.org/alien_soft/sms_service/pkg/logger"
	"bitbucket.org/alien_soft/sms_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SmsPaymentsService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewSmsPaymentsService(db *sqlx.DB, log l.Logger) *SmsPaymentsService {
	return &SmsPaymentsService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *SmsPaymentsService) Get(ctx context.Context, req *pb.SmsPaymentsId) (*pb.SmsPayments, error) {
	smsPayments, err := s.storage.SmsPayment().Get(req.Id)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting sms payments, Not found", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting sms payments", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return smsPayments, nil
}

func (s *SmsPaymentsService) Create(ctx context.Context, req *pb.SmsPayments) (*pb.SmsPaymentsId, error) {
	smsPayments, err := s.storage.SmsPayment().Create(req)
	if err == sql.ErrNoRows {
		s.logger.Error("Error while getting sms payments, Not found", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.NotFound, "Not found")
	} else if err != nil {
		s.logger.Error("Error while getting sms payments", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return smsPayments, nil
}

func (s *SmsPaymentsService) GetAll(ctx context.Context, req *pb.GetAllSmsPaymentsRequest) (*pb.GetAllSmsPaymentsResponse, error) {
	smsPayments, count, err := s.storage.SmsPayment().GetAll(req.GetShipperId(), int(req.GetPage()), int(req.GetLimit()))

	if err != nil {
		s.logger.Error("Error while getting sms payments, Not found", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetAllSmsPaymentsResponse{
		Smspayments: smsPayments,
		Count:       int64(count),
	}, nil
}
