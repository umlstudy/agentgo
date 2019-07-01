module sejong.asia/serverMonitor/agent

go 1.12

replace sejong.asia/serverMonitor/common => ../common

require (
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4
	github.com/stretchr/testify v1.3.0 // indirect
	sejong.asia/serverMonitor/common v0.0.0-00010101000000-000000000000
)
