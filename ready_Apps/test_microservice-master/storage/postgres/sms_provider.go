package postgres

import (
	"database/sql"
	pb "genproto/sms_service"
	"time"

	"bitbucket.org/alien_soft/sms_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type smsProviderRepo struct {
	db *sqlx.DB
}

func NewSmsProviderRepo(db *sqlx.DB) repo.SmsProviderStorageI {
	return &smsProviderRepo{db: db}
}

func (cm *smsProviderRepo) Get(shipperId string) (*pb.SmsProvider, error) {
	var (
		smsProvider pb.SmsProvider
		createdAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
	)

	row := cm.db.QueryRow(`
		SELECT  
			shipper_id,
			name,
			login,
			password,
			created_at
		FROM sms_providers
		WHERE shipper_id = $1`,
		shipperId,
	)

	err := row.Scan(
		&smsProvider.ShipperId,
		&smsProvider.Name,
		&smsProvider.Login,
		&smsProvider.Password,
		&createdAt,
	)

	if err != nil {
		return nil, err
	}

	smsProvider.CreatedAt = createdAt.Format(layoutDate)

	return &smsProvider, nil
}

func (cm *smsProviderRepo) Create(smsProvider *pb.SmsProvider) error {
	query :=
		`INSERT INTO sms_providers
			(
				shipper_id,
				name,
				login,
				password
			)
		VALUES ($1, $2, $3, $4)`

	_, err := cm.db.Exec(
		query,
		smsProvider.GetShipperId(),
		smsProvider.GetName(),
		smsProvider.GetLogin(),
		smsProvider.GetPassword(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (cm *smsProviderRepo) Delete(shipperId string) error {
	result, err := cm.db.Exec(`DELETE FROM sms_providers WHERE shipper_id=$1`, shipperId)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
