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
		return "User is offline."
	case OnlineStatus:
		return "User is oneline."
	case BusyStatus:
		return "User is busy."
	case LeaveTemporarily:
		return "User temporarily left."
	default:
		return "Illegal status."
	}
}
