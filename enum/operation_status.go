package enum

type OperationStatus uint16

const (
	InvalidOperation     OperationStatus = 0
	SuccessfulOpration   OperationStatus = 1
	FailedOpration       OperationStatus = 2
	UnauthorizedOpration OperationStatus = 3
)

// To convert the status to the corresponding string.
func (o OperationStatus) String() string {
	switch o {
	case InvalidOperation:
		return "Invalid Operation."
	case SuccessfulOpration:
		return "Successful operation."
	case FailedOpration:
		return "operation failed."
	case UnauthorizedOpration:
		return "Illegal operation."
	default:
		return "Illegal Status."
	}
}
