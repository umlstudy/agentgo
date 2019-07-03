package common

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func CreateServerInfo(pss []ProcessStatus, procNameParts []string, warningConditionMap map[string]WarningCondition) (*ServerInfo, error) {

	resourceStatuss := []ResourceStatus{}

	// CPU
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return nil, errors.Wrap(err, "cpu info read failed")
	}
	cpuSum := float64(0)
	for _, per := range percentage {
		cpuSum = cpuSum + per
	}
	cpuAvg := uint32(cpuSum / float64(len(percentage)))
	wc := warningConditionMap["cpu"]
	resourceStatuss = append(resourceStatuss, ResourceStatus{AbstractStatus{"cpu", "cpu", NORMAL, wc}, 1, 100, cpuAvg})

	// 메모리
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.Wrap(err, "memory info read failed")
	}
	wc = warningConditionMap["mem"]
	resourceStatuss = append(resourceStatuss, ResourceStatus{AbstractStatus{"mem", "mem", NORMAL, wc}, 1, 100, uint32(vmStat.UsedPercent)})

	// 파티션
	ptns, err := disk.Partitions(false)
	if err != nil {
		return nil, errors.Wrap(err, "partition info read failed")
	}
	for i, ptn := range ptns {
		diskStat, err := disk.Usage(ptn.Mountpoint)
		if err != nil {
			return nil, errors.Wrap(err, "disk info read failed")
		}
		wc = warningConditionMap[fmt.Sprint("disk%d", i)]
		resourceStatuss = append(resourceStatuss, ResourceStatus{AbstractStatus{diskStat.Path, diskStat.Path, NORMAL, wc}, 1, 100, uint32(diskStat.UsedPercent)})
	}

	// 프로세스
	pss, err = CheckAliveProcessStatuses(pss, procNameParts, warningConditionMap)
	if err != nil {
		return nil, errors.Wrap(err, "process info read failed")
	}

	// 머신정보
	hostStat, err := host.Info()
	if err != nil {
		return nil, errors.Wrap(err, "host info read failed")
	}

	wc = warningConditionMap["host"]
	serverInfo := &ServerInfo{AbstractStatus{hostStat.Hostname, fmt.Sprintf("%s(%s)", hostStat.Hostname, hostStat.Platform), NORMAL, wc}, resourceStatuss, pss}

	return serverInfo, nil
}
