package repo

import (
	pb "genproto/sms_service"
)

type SmsPaymentsStorageI interface {
	Create(smsPayment *pb.SmsPayments) (*pb.SmsPaymentsId, error)
	Get(id string) (*pb.SmsPayments, error)
	GetAll(shipperId string, page int, limit int) (smsPayment []*pb.SmsPayments, count int, er error)
}
