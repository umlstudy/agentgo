package common

// 1. ConditionType start
type ConditionType uint32

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

func (c ConditionType) GetLabel() string {
	return conditionTypeLabels[uint32(c-1)%uint32(len(conditionTypeLabels))]
}

// 1. ConditionType End

// 2. ConditionType start
type WarningLevel uint32

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

func (w WarningLevel) GetLabel() string {
	return warningLevelLabels[uint32(w-1)%uint32(len(warningLevelLabels))]
}

// 2. ConditionType End
