package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/w32"
)

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
		os.Exit(-1)
	}
}

func main() {

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
