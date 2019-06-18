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
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetLogicalDrives = modkernel32.NewProc("GetLogicalDrives")
)

func GetLogicalDrives() uint32 {
	ret, _, _ := procGetLogicalDrives.Call()

	return uint32(ret)
}

func dealwithErr2(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
		panic(err)
	}
}

// https://mingrammer.com/gobyexample/

///////////////////////

type Member struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type UsageStat2 struct {
	Path        string  `json:"path"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func convertToUsageStat2(us *disk.UsageStat) *UsageStat2 {
	us2 := &UsageStat2{
		Path:        us.Path,
		Total:       uint64(us.Total / 1024 / 1024 / 1024),
		Free:        uint64(us.Free / 1024 / 1024 / 1024),
		Used:        uint64(us.Used / 1024 / 1024 / 1024),
		UsedPercent: us.UsedPercent,
	}
	return us2
}

func Send(us *UsageStat2) {
	url := "https://restapi3.docs.apiary.io/notes"
	fmt.Println("URL:>", url)

	jsonBytes, err := json.Marshal(us)
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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {

	drvs := w32.GetLogicalDrives()

	validDriveLetters := []byte{}
	for x := 2; x < 10; x++ {
		mask := uint32(drvs) >> uint32(x)
		if (1 & mask) == uint32(1) {
			drvLetter := x + 65
			validDriveLetters = append(validDriveLetters, byte(x+65))

			diskStat, err := disk.Usage(fmt.Sprintf("%c:/", drvLetter))
			dealwithErr2(err)

			total := strconv.FormatUint(diskStat.Total/1024/1024/1024, 10)
			used := strconv.FormatUint(diskStat.Used/1024/1024/1024, 10)
			free := strconv.FormatUint(diskStat.Free/1024/1024/1024, 10)
			percentDiskSpaceUsage := strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

			fmt.Printf("%c Drive => Total:%sG, Used:%sG, Free:%sG(%s%%)\n", drvLetter, total, used, free, percentDiskSpaceUsage)
		} else {
		}
	}

	i := 0
	for true {
		i++
		time.Sleep(2 * time.Second)

		for _, driveLetter := range validDriveLetters {
			diskStat, err := disk.Usage(fmt.Sprintf("%c:/", driveLetter))
			dealwithErr2(err)

			diskStat2 := convertToUsageStat2(diskStat)
			Send(diskStat2)

			total := strconv.FormatUint(diskStat2.Total, 10)
			used := strconv.FormatUint(diskStat2.Used, 10)
			free := strconv.FormatUint(diskStat2.Free, 10)
			percentDiskSpaceUsage := strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

			fmt.Printf("%c Drive => Total:%sG, Used:%sG, Free:%sG(%s%%)\n", driveLetter, total, used, free, percentDiskSpaceUsage)
		}

		fmt.Printf("%d\n", i)
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
			dealwithErr2(err)

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
