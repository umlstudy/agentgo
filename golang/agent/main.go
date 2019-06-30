package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"sejong.asia/serverMonitor/common"
)

// ResourceStatus is ResourceStatus
type ResourceStatus = common.ResourceStatus

// ServerInfo is ServerInfo
type ServerInfo = common.ServerInfo

func send(si ServerInfo, url string) bool {
	jsonBytes, err := json.Marshal(si)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	} else {
		return false
	}
}

func createServerInfo() ServerInfo {
	resourceStatuss := []ResourceStatus{}

	// CPU
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	resourceStatuss = append(resourceStatuss, ResourceStatus{"mem", 1, 100, "mem", uint32(vmStat.UsedPercent)})

	// 파티션
	ptns, err := disk.Partitions(false)
	if err != nil {
		panic(err)
	}
	for _, ptn := range ptns {
		diskStat, err := disk.Usage(ptn.Device)
		if err != nil {
			panic(err)
		}
		resourceStatuss = append(resourceStatuss, ResourceStatus{diskStat.Path, 1, 100, diskStat.Path, uint32(diskStat.UsedPercent)})
	}

	// 머신정보
	hostStat, err := host.Info()
	if err != nil {
		panic(err)
	}

	serverInfo := ServerInfo{hostStat.Hostname, fmt.Sprintf("%s(%s)", hostStat.Hostname, hostStat.Platform), resourceStatuss}

	return serverInfo
}

const urlFormat string = "http://%s:%d/recvServerInfo"

func main() {

	host := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")

	flag.Parse()

	url := fmt.Sprintf(urlFormat, *host, *port)

	i := 0
	for true {
		i++
		time.Sleep(1 * time.Second)

		si := createServerInfo()
		isOk := send(si, url)

		if isOk {
			fmt.Printf(".")
		} else {
			panic(fmt.Sprintf("ServerMonitory Gateway is not running.(%s)", url))
		}
		if i%80 == 0 {
			fmt.Printf("\n")
		}
	}
}
