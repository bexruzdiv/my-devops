package storage

import (
	"bitbucket.org/alien_soft/sms_service/storage/postgres"
	"bitbucket.org/alien_soft/sms_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

// StorageI ...
type StorageI interface {
	SendSms() repo.SmsStorageI
	SmsProvider() repo.SmsProviderStorageI
	SmsPayment() repo.SmsPaymentsStorageI
}

type storagePg struct {
	db              *sqlx.DB
	smsRepo         repo.SmsStorageI
	smsProviderRepo repo.SmsProviderStorageI
	smsPaymentRepo  repo.SmsPaymentsStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:              db,
		smsRepo:         postgres.NewSmsRepo(db),
		smsProviderRepo: postgres.NewSmsProviderRepo(db),
		smsPaymentRepo:  postgres.NewPaymentsRepo(db),
	}
}

func (s storagePg) SendSms() repo.SmsStorageI {
	return s.smsRepo
}

func (s storagePg) SmsProvider() repo.SmsProviderStorageI {
	return s.smsProviderRepo
}

func (s storagePg) SmsPayment() repo.SmsPaymentsStorageI {
	return s.smsPaymentRepo
}
