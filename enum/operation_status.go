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
		return "invalid operation"
	case SuccessfulOpration:
		return "successful operation"
	case FailedOpration:
		return "operation failed"
	case UnauthorizedOpration:
		return "illegal operation"
	default:
		return "illegal status"
	}
}
