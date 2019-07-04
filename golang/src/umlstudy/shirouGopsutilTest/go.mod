module sejong.asia/serverMonitor/gotest

go 1.12

replace sejong.asia/serverMonitor/common => ../common

require (
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/stretchr/testify v1.3.0 // indirect
	sejong.asia/serverMonitor/common v0.0.0-00010101000000-000000000000
)
