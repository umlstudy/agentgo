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

type AgentSettings struct {
	ProcNameParts       []string                           `json:"procNameParts"`
	WarningConditionMap map[string]common.WarningCondition `json:"warningConditionMap"`
}

func readJson(fileName string) (*AgentSettings, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var as = AgentSettings{}
	err = json.Unmarshal([]byte(byteValue), &as)
	if err != nil {
		return nil, err
	}

	return &as, nil
}

func main() {

	host := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()

	as, err := readJson("agentSettings.json")
	if err != nil {
		panic(err)
	}
	pss, err := common.FindMatchedPids(as.ProcNameParts, as.WarningConditionMap)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(urlFormat, *host, *port)
	i := 0
	for true {

		time.Sleep(1 * time.Second)

		si, err := common.CreateServerInfo(pss, as.ProcNameParts, as.WarningConditionMap)
		if err != nil {
			panic(err)
		}

		err = common.SendPostWithJson(si, url)
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
