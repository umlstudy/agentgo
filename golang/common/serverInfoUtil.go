package common

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/pkg/errors"
)

func CreateServerInfo(pss []ProcessStatus, procNameParts []string) (*ServerInfo, error) {

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
	resourceStatuss = append(resourceStatuss, ResourceStatus{"cpu", 1, 100, "cpu", cpuAvg})

	// 메모리
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.Wrap(err, "memory info read failed")
	}
	resourceStatuss = append(resourceStatuss, ResourceStatus{"mem", 1, 100, "mem", uint32(vmStat.UsedPercent)})

	// 파티션
	ptns, err := disk.Partitions(false)
	if err != nil {
		return nil, errors.Wrap(err, "partition info read failed")
	}
	for _, ptn := range ptns {
		diskStat, err := disk.Usage(ptn.Mountpoint)
		if err != nil {
			return nil, errors.Wrap(err, "disk info read failed")
		}
		resourceStatuss = append(resourceStatuss, ResourceStatus{diskStat.Path, 1, 100, diskStat.Path, uint32(diskStat.UsedPercent)})
	}

	// 프로세스
	pss, err = CheckAliveProcessStatuses(pss, procNameParts)
	if err != nil {
		return nil, errors.Wrap(err, "process info read failed")
	}

	// 머신정보
	hostStat, err := host.Info()
	if err != nil {
		return nil, errors.Wrap(err, "host info read failed")
	}

	serverInfo := &ServerInfo{hostStat.Hostname, fmt.Sprintf("%s(%s)", hostStat.Hostname, hostStat.Platform), resourceStatuss, pss}

	return serverInfo, nil
}
