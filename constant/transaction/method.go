package transaction

type Method int

const (
	TransactionMethodCreditCard Method = iota + 1
	TransactionMethodPaypal
	TransactionMethodBankTransfer
	TransactionMethodCash
)

var mapTransactionMethod = map[Method]string{
	TransactionMethodCreditCard:   "CREDIT CARD",
	TransactionMethodPaypal:       "PAYPAL",
	TransactionMethodBankTransfer: "BANK TRANSFER",
	TransactionMethodCash:         "CASH",
}

func (m Method) Enum() string {
	if val, ok := mapTransactionMethod[m]; ok {
		return val
	}

	return "UNKNOWN"
}

func (m Method) IsValid() bool {
	if _, ok := mapTransactionMethod[m]; ok {
		return true
	}

	return false
}
