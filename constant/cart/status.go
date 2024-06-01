package cart

type Status int

const (
	CartStatusActive Status = iota + 1
	CartStatusCompleted
	CartStatusCancelled
)

var mapCartStatus = map[Status]string{
	CartStatusActive:    "Active",
	CartStatusCompleted: "Completed",
	CartStatusCancelled: "Cancelled",
}

func (s Status) Enum() string {
	if val, ok := mapCartStatus[s]; ok {
		return val
	}

	return "Unknown"
}