package transaction

type Status int

const (
	TransactionStatusInprogress Status = iota + 1
	TransactionStatusSuccess
	TransactionStatusFailed
)

var mapTransactionStatus = map[Status]string{
	TransactionStatusInprogress: "IN PROGRESS",
	TransactionStatusSuccess:    "SUCCESS",
	TransactionStatusFailed:     "FAILED",
}

func (s Status) Enum() string {
	if val, ok := mapTransactionStatus[s]; ok {
		return val
	}

	return "UNKNOWN"
}