package postgres

import (
	pb "genproto/sms_service"
	"time"

	"bitbucket.org/alien_soft/sms_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type smsPaymentsRepo struct {
	db *sqlx.DB
}

func NewPaymentsRepo(db *sqlx.DB) repo.SmsPaymentsStorageI {
	return &smsPaymentsRepo{db: db}
}

func (cm *smsPaymentsRepo) Create(smsPayments *pb.SmsPayments) (*pb.SmsPaymentsId, error) {
	var (
		resp pb.SmsPaymentsId
	)

	query := `
		INSERT INTO sms_payments
			(
				id,
				shipper_id,
				order_id,
				ext_order_id,
				phone,
				operator_id,
				operator_name,
				payment_type,
				order_amount
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	_, err = cm.db.Exec(
		query,
		id,
		smsPayments.GetShipperId(),
		smsPayments.GetOrderId(),
		smsPayments.GetExtOrderId(),
		smsPayments.GetPhone(),
		smsPayments.GetOperatorId(),
		smsPayments.GetOperatorName(),
		smsPayments.GetPaymentType(),
		smsPayments.GetOrderAmount(),
	)

	if err != nil {
		return nil, err
	}

	resp.Id = id.String()

	return &resp, nil
}

func (cm *smsPaymentsRepo) Get(id string) (*pb.SmsPayments, error) {
	var (
		smsPayments pb.SmsPayments
		createdAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
	)

	row := cm.db.QueryRow(`
		SELECT
			id,
			shipper_id,
			order_id,
			ext_order_id,
			phone,
			operator_id,
			operator_name,
			payment_type,
			order_amount,
			created_at
		FROM sms_payments
		WHERE id = $1`,
		id,
	)

	err := row.Scan(
		&smsPayments.Id,
		&smsPayments.ShipperId,
		&smsPayments.OrderId,
		&smsPayments.ExtOrderId,
		&smsPayments.Phone,
		&smsPayments.OperatorId,
		&smsPayments.OperatorName,
		&smsPayments.PaymentType,
		&smsPayments.OrderAmount,
		&createdAt,
	)

	if err != nil {
		return nil, err
	}

	smsPayments.CreatedAt = createdAt.Format(layoutDate)

	return &smsPayments, nil
}

func (cm *smsPaymentsRepo) GetAll(shipperID string, page int, limit int) ([]*pb.SmsPayments, int, error) {

	var (
		createdAt      time.Time
		smsPaymentsArr []*pb.SmsPayments
		count          int64
		layoutDate     string = "2006-01-02 15:04:05"
	)

	offset := (page - 1) * limit

	rows, err := cm.db.Query(`
		SELECT
			id,
			shipper_id,
			order_id,
			ext_order_id,
			phone,
			operator_id,
			operator_name,
			payment_type,
			order_amount,
			created_at
		FROM sms_payments
		WHERE shipper_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`,
		shipperID,
		limit,
		offset,
	)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var smsPayments pb.SmsPayments

		err := rows.Scan(
			&smsPayments.Id,
			&smsPayments.ShipperId,
			&smsPayments.OrderId,
			&smsPayments.ExtOrderId,
			&smsPayments.Phone,
			&smsPayments.OperatorId,
			&smsPayments.OperatorName,
			&smsPayments.PaymentType,
			&smsPayments.OrderAmount,
			&createdAt,
		)

		if err != nil {
			return nil, 0, err
		}

		smsPayments.CreatedAt = createdAt.Format(layoutDate)
		smsPaymentsArr = append(smsPaymentsArr, &smsPayments)
	}

	row := cm.db.QueryRow(`
		SELECT count(1)
		FROM sms_payments
		WHERE shipper_id = $1`,
		shipperID,
	)

	err = row.Scan(
		&count,
	)

	if err != nil {
		return nil, 0, err
	}

	return smsPaymentsArr, int(count), nil
}
