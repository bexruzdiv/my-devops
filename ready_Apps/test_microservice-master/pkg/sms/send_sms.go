package mail

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"bitbucket.org/alien_soft/sms_service/pkg/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //db driver
	"github.com/sfreiberg/gotwilio"

	"bitbucket.org/alien_soft/sms_service/config"
	"bitbucket.org/alien_soft/sms_service/storage/postgres"
	rp "bitbucket.org/alien_soft/sms_service/storage/repo"
)

//Daemon ...
type Daemon struct {
	Conf *config.Config
	DB   *sqlx.DB
}

var (
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

// Init initializes Deamon
func (dmn *Daemon) Init() {
	repo := postgres.NewSmsRepo(dmn.DB)

	c := make(chan string)

	isShuttingDown := false
	started := false

OUTTER:
	for {
		if isShuttingDown {
			break
		}
		select {
		case s, ok := <-c:

			if ok {
				fmt.Printf("\nservice is going down...\n")
				fmt.Println(fmt.Sprintf("\nReceived signal %x\n", s))
				isShuttingDown = true
				c = nil
				continue OUTTER
			}
		default:
			var smss []rp.Sms
			smss, err := repo.GetNotSent()
			if err != nil {
				log.Println("Error while checking status: ", err)
			}
			for _, sms := range smss {

				err := dmn.sendSms(sms.ShipperId, sms.Text, sms.PhoneNumber)
				if err != nil {
					log.Println("error while sending sms: ", err)

					err := repo.IncrementSendCount(sms.ID)
					log.Println("error while incrementing sms send trying count ", err)
				} else {
					err = repo.MakeSent(sms.ID)
					if err != nil {
						log.Println("error while updating status")
					}
				}
			}

			if !started {
				fmt.Print("\n\nSuccessfully\n\n")
				started = true
			}
			t := time.Now()
			fmt.Println(fmt.Sprint("Service is running...\n", t.Format(time.RFC3339)))
			time.Sleep(15000 * time.Millisecond)
		}
	}
	fmt.Print("\n\nThis service has finally shutted down\n\n")
}

// Send Mail...
func (dmn *Daemon) sendSms(shipperId, text, phoneNumber string) error {
	phone := phoneNumber[1:]
	fmt.Println(phone)
	if dmn.Conf.SMSSender == "play_mobile" {
		return sendWithPlayMobile(dmn, shipperId, text, phone)
	}

	return sendWithTwilio(dmn, text, phone)
}

func sendWithTwilio(dmn *Daemon, text, phoneNumber string) error {
	accountSid := dmn.Conf.AccountSid
	authToken := dmn.Conf.AuthToken

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := dmn.Conf.PhoneNumber
	to := phoneNumber
	message := text

	_, exception, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return err
	} else if exception != nil {
		return exception
	}

	return nil
}

func sendWithPlayMobile(dmn *Daemon, shipperId, text, phoneNumber string) error {
	repo := postgres.NewSmsProviderRepo(dmn.DB)
	code, err := GenerateCode(15)

	if err != nil {
		return err
	}

	var body models.Body
	client := http.Client{}

	message := models.Message{
		Recipient: phoneNumber,
		MessageID: fmt.Sprintf("%s%s", "oql", code),
		SMS: models.SMS{
			Originator: dmn.Conf.PlayMobileOriginator,
			Content: models.Content{
				Text: text,
			},
		},
	}

	body.Messages = append(body.Messages, message)

	requestBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", dmn.Conf.PlayMobileUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	smsProvider, err := repo.Get(shipperId)
	if err != nil {
		return err
	}

	request.SetBasicAuth(smsProvider.Login, smsProvider.Password)

	res, err := client.Do(request)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("error while sending sms")
	}

	return nil
}

func GenerateCode(max int) (string, error) {
	b := make([]byte, max)

	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b), nil
}
