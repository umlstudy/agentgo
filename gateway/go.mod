module github.com/umlstudy/serverMonitor/gateway

go 1.12

replace github.com/umlstudy/serverMonitor/common => ../common

require (
	github.com/fatih/color v1.7.0
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
	github.com/takama/daemon v0.11.0
	github.com/umlstudy/serverMonitor/common v0.0.0-00010101000000-000000000000
)
