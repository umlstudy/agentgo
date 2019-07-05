package common

// 1. ConditionType start

// ConditionType is ConditionType
type ConditionType uint32

// Less is Less
const (
	Less ConditionType = 1 + iota
	LessOrEqual
	Equal
	GreaterOrEqual
	Greater
)

var conditionTypeLabels = [...]string{
	"<",
	"<=",
	"=",
	">=",
	">",
}

// GetLabel is GetLabel
func (c ConditionType) GetLabel() string {
	return conditionTypeLabels[uint32(c-1)%uint32(len(conditionTypeLabels))]
}

// 1. ConditionType End

// 2. WarningLevel start

// WarningLevel is WarningLevel
type WarningLevel uint32

// WarningLevel is WarningLevel
const (
	NORMAL WarningLevel = 1 + iota
	WARNING
	ERROR
)

var warningLevelLabels = []string{
	"NORMAL",
	"WARNING",
	"ERROR",
}

// GetLabel is GetLabel
func (w WarningLevel) GetLabel() string {
	return warningLevelLabels[uint32(w-1)%uint32(len(warningLevelLabels))]
}

// 2. WarningLevel End
