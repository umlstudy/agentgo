package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/umlstudy/serverMonitor/common"
)

func determineWarningLevelA(value uint32, warningLevelChangeConditionMap map[string]common.WarningLevelChangeCondition) common.WarningLevel {
	wlcc := warningLevelChangeConditionMap[common.ERROR.GetLabel()]
	if determineWarningLevelB(value, wlcc) {
		return common.ERROR
	}
	wlcc = warningLevelChangeConditionMap[common.WARNING.GetLabel()]
	if determineWarningLevelB(value, wlcc) {
		return common.WARNING
	}
	return common.NORMAL
}

func determineWarningLevelB(value uint32, wlcc common.WarningLevelChangeCondition) bool {
	if wlcc.ConditionType < 1 {
		// 경고레벨 변경 컨디션이 설정 안되어 있는 경우
		return false
	}
	switch wlcc.ConditionType {
	case common.Less:
		if value < wlcc.Value {
			return true
		}
		break
	case common.LessOrEqual:
		if value <= wlcc.Value {
			return true
		}
		break
	case common.Equal:
		if value == wlcc.Value {
			return true
		}
		break
	case common.GreaterOrEqual:
		if value >= wlcc.Value {
			return true
		}
		break
	case common.Greater:
		if value > wlcc.Value {
			return true
		}
		break
	}

	return false
}

func createServerInfo(pss []common.ProcessStatus, procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) (*common.ServerInfo, error) {

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
	resourceStatuss = append(resourceStatuss, common.ResourceStatus{AbstractStatus: common.AbstractStatus{ID: "cpu", Name: "cpu", WarningLevel: wl, AlarmCondition: ac}, Min: 1, Max: 100, Value: cpuAvg})

	// 메모리
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, errors.Wrap(err, "memory info read failed")
	}
	memUsed := uint32(vmStat.UsedPercent)
	acwwlcc = alarmConditionWithWarningLevelChangeConditionMap["mem"]
	ac = acwwlcc.AlarmCondition
	wl = determineWarningLevelA(memUsed, acwwlcc.WarningLevelChangeConditionMap)
	resourceStatuss = append(resourceStatuss, common.ResourceStatus{AbstractStatus: common.AbstractStatus{ID: "mem", Name: "mem", WarningLevel: wl, AlarmCondition: ac}, Min: 1, Max: 100, Value: memUsed})

	// 파티션
	ptns, err := disk.Partitions(false)
	if err != nil {
		return nil, errors.Wrap(err, "partition info read failed")
	}
	for i, ptn := range ptns {
		if strings.Contains(ptn.Mountpoint, "/var/lib/docker") {
			continue
		}
		diskStat, err := disk.Usage(ptn.Mountpoint)
		if err != nil {
			return nil, errors.Wrap(err, "disk info read failed")
		}
		diskUsed := uint32(diskStat.UsedPercent)
		acwwlcc = alarmConditionWithWarningLevelChangeConditionMap[fmt.Sprintf("disk%d", i)]
		ac = acwwlcc.AlarmCondition
		wl = determineWarningLevelA(diskUsed, acwwlcc.WarningLevelChangeConditionMap)
		resourceStatuss = append(resourceStatuss, common.ResourceStatus{AbstractStatus: common.AbstractStatus{ID: diskStat.Path, Name: diskStat.Path, WarningLevel: wl, AlarmCondition: ac}, Min: 1, Max: 100, Value: diskUsed})
	}

	// 프로세스
	pss, err = checkAliveProcessStatuses(pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
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
	serverInfo := &common.ServerInfo{ID: hostStat.Hostname, Name: fmt.Sprintf("%s(%s)", hostStat.Hostname, hostStat.Platform), AlarmCondition: ac, IsRunning: true, ResourceStatuses: resourceStatuss, ProcessStatuses: pss}

	return serverInfo, nil
}
