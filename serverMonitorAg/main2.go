package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/w32"
)

//import myTypes "github.com/shirou/gopsutil/disk/UsageStat"

var (
	modkernel32          = syscall.NewLazyDLL("kernel32.dll")
	procGetLogicalDrives = modkernel32.NewProc("GetLogicalDrives")
)

func GetLogicalDrives() uint32 {
	ret, _, _ := procGetLogicalDrives.Call()
	return uint32(ret)
}

// https://mingrammer.com/gobyexample/

///////////////////////

type Member struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

// type UsageStat2 struct {
// 	Path        string  `json:"path"`
// 	Fstype      string  `json:"fstype"`
// 	Total       uint64  `json:"total"`
// 	Free        uint64  `json:"free"`
// 	Used        uint64  `json:"used"`
// 	UsedPercent float64 `json:"usedPercent"`
// }

type ResourceStatus struct {
	Id    string `json:"id"`
	Min   uint32 `json:"min"`
	Max   uint32 `json:"max"`
	Name  string `json:"name"`
	Value uint32 `json:"value"`
}
type ServerInfo struct {
	Id               string           `json:"id"`
	Name             string           `json:"name"`
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
}

func convertToResourceStatus(us *disk.UsageStat, id string) ResourceStatus {
	// Path:        us.Path,
	// Total:       uint64(us.Total / 1024 / 1024 / 1024),
	// Free:        uint64(us.Free / 1024 / 1024 / 1024),
	// Used:        uint64(us.Used / 1024 / 1024 / 1024),
	// UsedPercent: us.UsedPercent,
	var val = uint32(us.UsedPercent)
	rs := ResourceStatus{
		Id:    id,
		Name:  us.Path,
		Min:   0,
		Max:   100,
		Value: val}
	return rs
}

const url string = "http://localhost:8080/recvServerInfo"

func Send(si ServerInfo) bool {
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
	// respStatCode, err := strconv.Atoi(resp.StatusCode)
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
}

func getValidDriveLetters() []byte {
	drvs := w32.GetLogicalDrives()

	validDriveLetters := []byte{}
	for x := 2; x < 10; x++ {
		mask := uint32(drvs) >> uint32(x)
		if (1 & mask) == uint32(1) {
			drvLetter := x + 65
			validDriveLetters = append(validDriveLetters, byte(drvLetter))
		}
	}

	return validDriveLetters
}

func main() {

	validDriveLetters := getValidDriveLetters()

	i := 0
	for true {
		i++
		time.Sleep(1 * time.Second)

		var rss []ResourceStatus
		for _, driveLetter := range validDriveLetters {
			driveLetterStr := fmt.Sprintf("%c:/", driveLetter)
			diskStat, err := disk.Usage(driveLetterStr)
			if err != nil {
				panic(err)
			}

			rs := convertToResourceStatus(diskStat, driveLetterStr)
			rss = append(rss, rs)
		}
		si := ServerInfo{"mysvr", "mysvr", rss}
		isOk := Send(si)

		if isOk {
			fmt.Printf(".")
		} else {
			fmt.Printf("x")
		}
		if i%80 == 0 {
			fmt.Printf("\n")
		}
	}
}

func main4() {

	// Go 데이타
	mem := Member{"Alex", 10, true}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)

	// JSON 디코딩
	var mem2 Member
	err = json.Unmarshal(jsonBytes, &mem2)
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem2.Name, mem2.Age, mem2.Active)
}

func main3() {
	url := "https://restapi3.docs.apiary.io/notes"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main2() {

	drvs := w32.GetLogicalDrives()
	//fmt.Printf(strconv.FormatUint(uint64(drvs), 2) + "\n")

	for x := 2; x < 10; x++ {
		mask := uint32(drvs) >> uint32(x)
		//fmt.Printf(strconv.FormatUint(uint64(mask), 2) + "\n")
		if (1 & mask) == uint32(1) {
			drvLetter := x + 65
			//fmt.Printf("%c drive is on.\n", drvLetter)

			diskStat, err := disk.Usage(fmt.Sprintf("%c:/", drvLetter))
			if err != nil {
				panic(err)
			}

			total := strconv.FormatUint(diskStat.Total/1024/1024/1024, 10)
			used := strconv.FormatUint(diskStat.Used/1024/1024/1024, 10)
			free := strconv.FormatUint(diskStat.Free/1024/1024/1024, 10)
			percentDiskSpaceUsage := strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

			fmt.Printf("%c Drive => Total:%sG, Used:%sG, Free:%sG(%s%%)\n", drvLetter, total, used, free, percentDiskSpaceUsage)
		} else {
		}
	}

	// html := ""
	// html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total/1024/1024/1024, 10) + "G bytes \n"
	// html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used/1024/1024/1024, 10) + "G bytes\n"
	// html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free/1024/1024/1024, 10) + "G bytes\n"
	// html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%"

	// fmt.Println(html)
}
