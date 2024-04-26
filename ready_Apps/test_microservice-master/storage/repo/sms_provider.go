package repo

import (
	pb "genproto/sms_service"
)

type SmsProviderStorageI interface {
	Get(shipperId string) (*pb.SmsProvider, error)
	Create(smsProvider *pb.SmsProvider) error
	Delete(shipperId string) error
}
