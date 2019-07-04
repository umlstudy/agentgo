package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"sejong.asia/serverMonitor/common"
)

func determineWarningLevelA(value uint32, warningLevelChangeConditionMap map[common.WarningLevel]common.WarningLevelChangeCondition) common.WarningLevel {
	wlcc := warningLevelChangeConditionMap[common.ERROR]
	if determineWarningLevelB(value, wlcc) == common.ERROR {
		return common.ERROR
	}
	wlcc = warningLevelChangeConditionMap[common.WARNING]
	if determineWarningLevelB(value, wlcc) == common.WARNING {
		return common.WARNING
	}
	return common.NORMAL
}

func determineWarningLevelB(value uint32, wlcc common.WarningLevelChangeCondition) common.WarningLevel {
	switch wlcc.ConditionType {
	case common.Less:
		if value < wlcc.Value {
			return wlcc.WarningLevel
		}
		break
	case common.LessOrEqual:
		if value <= wlcc.Value {
			return wlcc.WarningLevel
		}
		break
	case common.Equal:
		if value == wlcc.Value {
			return wlcc.WarningLevel
		}
		break
	case common.GreaterOrEqual:
		if value >= wlcc.Value {
			return wlcc.WarningLevel
		}
		break
	case common.Greater:
		if value > wlcc.Value {
			return wlcc.WarningLevel
		}
		break
	}

	return common.NORMAL
}

func CreateServerInfo(pss []common.ProcessStatus, procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) (*common.ServerInfo, error) {

	resourceStatuss := []common.ResourceStatus{}

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
	acwwlcc := alarmConditionWithWarningLevelChangeConditionMap["cpu"]
	ac := acwwlcc.AlarmCondition
	wl := determineWarningLevelA(cpuAvg, acwwlcc.WarningLevelChangeConditionMap)
	resourceStatuss = append(resourceStatuss, common.ResourceStatus{common.AbstractStatus{"cpu", "cpu", wl, ac}, 1, 100, cpuAvg})

	// 메모리
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.Wrap(err, "memory info read failed")
	}
	memUsed := uint32(vmStat.UsedPercent)
	acwwlcc = alarmConditionWithWarningLevelChangeConditionMap["mem"]
	ac = acwwlcc.AlarmCondition
	wl = determineWarningLevelA(memUsed, acwwlcc.WarningLevelChangeConditionMap)
	resourceStatuss = append(resourceStatuss, common.ResourceStatus{common.AbstractStatus{"mem", "mem", wl, ac}, 1, 100, memUsed})

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
		diskUsed := uint32(diskStat.UsedPercent)
		acwwlcc = alarmConditionWithWarningLevelChangeConditionMap[fmt.Sprintf("disk%d", i)]
		ac = acwwlcc.AlarmCondition
		wl = determineWarningLevelA(diskUsed, acwwlcc.WarningLevelChangeConditionMap)
		resourceStatuss = append(resourceStatuss, common.ResourceStatus{common.AbstractStatus{diskStat.Path, diskStat.Path, wl, ac}, 1, 100, diskUsed})
	}

	// 프로세스
	pss, err = CheckAliveProcessStatuses(pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
	if err != nil {
		return nil, errors.Wrap(err, "process info read failed")
	}

	// 머신정보
	hostStat, err := host.Info()
	if err != nil {
		return nil, errors.Wrap(err, "host info read failed")
	}

	acwwlcc = alarmConditionWithWarningLevelChangeConditionMap["host"]
	ac = acwwlcc.AlarmCondition
	serverInfo := &common.ServerInfo{common.AbstractStatus{hostStat.Hostname, fmt.Sprintf("%s(%s)", hostStat.Hostname, hostStat.Platform), common.NORMAL, ac}, resourceStatuss, pss}

	return serverInfo, nil
}
