package repo

//Sms ...
type Sms struct {
	ID          string
	ShipperId   string
	Text        string
	PhoneNumber string
	SendCount   int
}

// SmstorageI ...
type SmsStorageI interface {
	GetNotSent() ([]Sms, error)
	MakeSent(ID string) error
	IncrementSendCount(ID string) error
	Send(shipperId, text string, recipients []string) error
}
