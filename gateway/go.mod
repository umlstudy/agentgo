module github.com/umlstudy/serverMonitor/gateway

go 1.12

replace github.com/umlstudy/serverMonitor/common => ../common

require (
	github.com/fatih/color v1.7.0
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/umlstudy/serverMonitor/common v0.0.0-00010101000000-000000000000
)
