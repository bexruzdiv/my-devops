package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"bitbucket.org/alien_soft/sms_service/storage/repo"
)

type smsRepo struct {
	db *sqlx.DB
}

// NewSmsRepo ...
func NewSmsRepo(db *sqlx.DB) repo.SmsStorageI {
	return &smsRepo{db: db}
}

// GetNotSent ...
func (cm *smsRepo) GetNotSent() ([]repo.Sms, error) {
	var smss []repo.Sms
	rows, err := cm.db.Queryx(`
		SELECT  id, 
				shipper_id,
				phone_number, 
				text,
				send_count
		FROM sms_send
		WHERE sent_at IS NULL and send_count < 4`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var sms repo.Sms
		if err := rows.Scan(
			&sms.ID,
			&sms.ShipperId,
			&sms.PhoneNumber,
			&sms.Text,
			&sms.SendCount,
		); err != nil {
			return nil, err
		}
		smss = append(smss, sms)
	}
	return smss, err
}

//MakeSent ...
func (cm *smsRepo) MakeSent(ID string) error {
	var err error
	makesent := `UPDATE sms_send SET sent_at = CURRENT_TIMESTAMP where id = $1`
	cm.db.MustExec(makesent, ID)
	return err
}

func (cm *smsRepo) IncrementSendCount(ID string) error {
	var err error
	query := `UPDATE sms_send SET send_count = send_count + 1 where id = $1`
	cm.db.MustExec(query, ID)
	return err
}

func (cm *smsRepo) Send(shipperId, text string, recipients []string) error {
	for _, recipient := range recipients {
		sendID, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		_, err = cm.db.Exec(`
			INSERT INTO
			sms_send
			(
				id,
				shipper_id,
				text,
				phone_number
			)
			values($1, $2, $3, $4)`, sendID, shipperId, text, recipient)
		if err != nil {
			return err
		}
	}

	return nil
}
