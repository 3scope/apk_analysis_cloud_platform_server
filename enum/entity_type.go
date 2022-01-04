package enum

type EntityType uint16

const (
	UserEntity            EntityType = 0
	CaseEntity            EntityType = 1
	StaticAnalysisEntity  EntityType = 2
	DynamicAnalysisEntity EntityType = 3
	ReportEntity          EntityType = 4
	VideoEntity           EntityType = 5
)

// To convert the status to the corresponding string.
func (e EntityType) String() string {
	switch e {
	case UserEntity:
		return "user entity"
	case CaseEntity:
		return "case entity"
	case StaticAnalysisEntity:
		return "static analysis entity"
	case DynamicAnalysisEntity:
		return "dynamic analysis entity"
	case ReportEntity:
		return "report entity"
	case VideoEntity:
		return "video entity"
	default:
		return "illegal type"
	}
}
