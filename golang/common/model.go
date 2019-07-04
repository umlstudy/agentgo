package common

// WarningCondition is WarningCondition
type WarningCondition struct {
	SendAlarmOccuredAfter    uint32 `json:"sendAlarmOccuredAfter"`
	ResendAlarmLastSendAfter uint32 `json:"resendAlarmLastSendAfter"`
}

type AbstractStatus struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	WarningLevel WarningLevel `json:"warningLevel"`
	WarningCondition
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
	AbstractStatus
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
	ProcessStatuses  []ProcessStatus  `json:"processStatuses"`
}

// ProcessStatus is ProcessStatus
type ProcessStatus struct {
	AbstractStatus
	RealName string `json:"realName"`
	ProcId   int32  `json:"procId"`
}

// DefaultServerPort is DefaultServerPort
const DefaultServerPort = 7007
