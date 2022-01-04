package enum

type UserStatus uint16

const (
	OfflineStatus    UserStatus = 0
	OnlineStatus     UserStatus = 1
	BusyStatus       UserStatus = 2
	LeaveTemporarily UserStatus = 3
)

// To Convert the status to the corresponding string.
func (u UserStatus) String() string {
	switch u {
	case OfflineStatus:
		return "user is offline"
	case OnlineStatus:
		return "user is oneline"
	case BusyStatus:
		return "user is busy"
	case LeaveTemporarily:
		return "user temporarily left"
	default:
		return "illegal status"
	}
}
