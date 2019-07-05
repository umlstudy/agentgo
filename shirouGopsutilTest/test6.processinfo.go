package main

import (
	"fmt"
	"strings"

	process "github.com/shirou/gopsutil/process"
	common "github.com/umlstudy/serverMonitor/common"
)

type ProcessStatus = common.ProcessStatus

func printProcess() {
	//ret, err := process.NewProcess(int32(1))
	pids, err := process.Pids()
	if err != nil {
		panic(err)
	}
	if len(pids) == 0 {
		panic("could not get pids")
	}
	for index, pid := range pids {
		println(index, pid)
		proc, err := process.NewProcess(int32(pid))
		if err != nil {
			panic(err)
		}
		procName, err := proc.Name()
		if err != nil {
			panic(err)
		}
		if strings.Contains(procName, "host") {
			fmt.Println(procName)
		}
	}
}

func test6() {
	fmt.Printf("test start...\n")
	// printProcess()
	// procNameParts := []string{"taskh", "supper"}
	// pss, err := common.FindMatchedPids(procNameParts)
	// if err != nil {
	// 	panic(err)
	// }
	// pss, err = common.CheckAliveProcessStatuses(pss, procNameParts)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ps := range pss {
	// 	fmt.Printf("%s %d\n", ps.RealName, ps.ProcId)
	// }
	fmt.Printf("test end...\n")
}
