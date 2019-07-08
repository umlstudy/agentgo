package common

// AlarmCondition is AlarmCondition
type AlarmCondition struct {
	SendAlarmOccuredAfter    uint64 `json:"sendAlarmOccuredAfter"`
	ResendAlarmLastSendAfter uint64 `json:"resendAlarmLastSendAfter"`
}

// WarningLevelChangeCondition is WarningLevelChangeCondition
type WarningLevelChangeCondition struct {
	ConditionType ConditionType `json:"conditionType"`
	Value         uint32        `json:"value"`
}

// AlarmConditionWithWarningLevelChangeCondition is AlarmConditionWithWarningLevelChangeCondition
type AlarmConditionWithWarningLevelChangeCondition struct {
	AlarmCondition
	WarningLevelChangeConditionMap map[string]WarningLevelChangeCondition `json:"warningLevelChangeConditionMap"`
}

// AbstractStatus is AbstractStatus
type AbstractStatus struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	WarningLevel WarningLevel `json:"warningLevel"`
	AlarmCondition
}

// ResourceStatus is ResourceStatus
type ResourceStatus struct {
	AbstractStatus
	Min   uint32 `json:"min"`
	Max   uint32 `json:"max"`
	Value uint32 `json:"value"`
}

// ServerInfo is ServerInfo
type ServerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	AlarmCondition
	IsRunning        bool             `json:"isRunning"`
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
	ProcessStatuses  []ProcessStatus  `json:"processStatuses"`
}

// ProcessStatus is ProcessStatus
type ProcessStatus struct {
	AbstractStatus
	RealName string `json:"realName"`
	ProcID   int32  `json:"procId"`
}

// DefaultServerPort is DefaultServerPort
const DefaultServerPort = 7007
