package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"sejong.asia/serverMonitor/common"
)

const urlFormat string = "http://%s:%d/recvServerInfo"

type AgentProperties struct {
	ProcNameParts []string `json:"procNameParts"`
}

func readJson(fileName string) (*AgentProperties, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var ap = AgentProperties{}
	err = json.Unmarshal([]byte(byteValue), &ap)
	if err != nil {
		return nil, err
	}

	return &ap, nil
}

func main() {

	host := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()

	ap, err := readJson("setting.json")
	if err != nil {
		panic(err)
	}
	pss, err := common.FindMatchedPids(ap.ProcNameParts)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(urlFormat, *host, *port)
	i := 0
	for true {

		time.Sleep(1 * time.Second)

		si, err := common.CreateServerInfo(pss, ap.ProcNameParts)
		if err != nil {
			panic(err)
		}

		err = common.SendPostToJson(si, url)
		if err != nil {
			panic(err)
		}

		fmt.Printf(".")
		i++
		if i%80 == 0 {
			fmt.Printf("\n")
		}
	}
}
