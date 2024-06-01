package cart

type Status int

const (
	CartStatusActive Status = iota + 1
	CartStatusPending
	CartStatusCompleted
	CartStatusCancelled
)

var mapCartStatus = map[Status]string{
	CartStatusActive:    "ACTIVE",
	CartStatusPending:   "PENDING",
	CartStatusCompleted: "COMPLETED",
	CartStatusCancelled: "CANCELLED",
}

func (s Status) Enum() string {
	if val, ok := mapCartStatus[s]; ok {
		return val
	}

	return "UNKNOWN"
}
