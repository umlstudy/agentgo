package common

// ResourceStatus is ResourceStatus
type ResourceStatus struct {
	ID    string `json:"id"`
	Min   uint32 `json:"min"`
	Max   uint32 `json:"max"`
	Name  string `json:"name"`
	Value uint32 `json:"value"`
}

// ServerInfo is ServerInfo
type ServerInfo struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
}
