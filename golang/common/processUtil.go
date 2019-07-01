package common

import (
	"errors"
	"strings"

	process "github.com/shirou/gopsutil/process"
)

func FindMatchedPids(procNameParts []string) ([]ProcessStatus, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}
	if len(pids) == 0 {
		return nil, errors.New("could not get pids")
	}
	processStatuses := []ProcessStatus{}
	for _, pid := range pids {
		proc, err := process.NewProcess(int32(pid))
		if err != nil {
			return nil, err
		}
		procName, err := proc.Name()
		if err != nil {
			return nil, err
		}
		for _, procNamePart := range procNameParts {
			if strings.Contains(procName, procNamePart) {
				processStatuses = append(processStatuses, ProcessStatus{procNamePart, procNamePart, procName, pid})
				continue
			}
		}
	}

	// 일치하는 프로세스가 없을 경우
	for _, procNamePart := range procNameParts {
		found := false
		for _, ps := range processStatuses {
			if strings.Contains(ps.RealName, procNamePart) {
				found = true
				continue
			}
		}
		if !found {
			processStatuses = append(processStatuses, ProcessStatus{procNamePart, procNamePart, procNamePart, 0})
		}
	}

	return processStatuses, nil
}

func CheckAliveProcessStatuses(pss []ProcessStatus, procNameParts []string) ([]ProcessStatus, error) {
	if len(pss) < 1 {
		return nil, errors.New("empty process statuses")
	}

	var newPss []ProcessStatus = nil
	for _, ps := range pss {
		if ps.ProcId < 1 {
			newPss_, err := FindMatchedPids(procNameParts)
			if err != nil {
				return nil, err
			}
			newPss = newPss_
			break
		}
	}

	if newPss == nil {
		for _, ps := range pss {

			proc, err := process.NewProcess(int32(ps.ProcId))
			if err != nil {
				ps.ProcId = 0
			} else {
				procName, err := proc.Name()
				if err != nil {
					ps.ProcId = 0
				} else {
					if strings.Compare(ps.RealName, procName) != 0 {
						ps.ProcId = 0
					}
				}
			}
		}

		newPss = pss
	}

	return newPss, nil
}
