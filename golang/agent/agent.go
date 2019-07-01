package main

import (
	"flag"
	"fmt"
	"time"

	"sejong.asia/serverMonitor/common"
)

const urlFormat string = "http://%s:%d/recvServerInfo"

func main() {

	host := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")

	flag.Parse()

	url := fmt.Sprintf(urlFormat, *host, *port)

	procNameParts := []string{"java"}
	pss, err := common.FindMatchedPids(procNameParts)
	if err != nil {
		panic(fmt.Errorf("ServerMonitory Gateway FindMatchedPids error.(%s, %s)", err, url))
	}

	i := 0
	for true {
		i++
		time.Sleep(1 * time.Second)
		si, err := common.CreateServerInfo(pss, procNameParts)
		if err != nil {
			panic(fmt.Errorf("ServerMonitory Gateway error.(%s, %s)", err, url))
		}

		err = common.SendPostToJson(si, url)
		if err != nil {
			panic(fmt.Errorf("ServerMonitory Gateway is not running.(%s, %s)", err, url))
		}
		fmt.Printf(".")
		if i%80 == 0 {
			fmt.Printf("\n")
		}
	}
}
