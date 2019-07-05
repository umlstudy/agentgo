package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-errors/errors"
	"github.com/umlstudy/serverMonitor/common"
)

const urlFormat string = "http://%s:%d/recvServerInfo"

// https://dangerous-animal141.hatenablog.com/entry/2017/01/19/004650
// enum -> json

// AgentSettings is AgentSettings
type AgentSettings struct {
	ProcNameParts                                    []string                                                        `json:"procNameParts"`
	AlarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition `json:"alarmConditionWithWarningLevelChangeConditionMap"`
}

func readJSON(fileName string) (*AgentSettings, error) {

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

	as, err := readJSON("agentSettings.json")
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		panic(err)
	}
	pss, err := findMatchedPids(as.ProcNameParts, as.AlarmConditionWithWarningLevelChangeConditionMap)
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		panic(err)
	}

	url := fmt.Sprintf(urlFormat, *host, *port)
	i := 0
	for true {

		time.Sleep(5 * time.Second)

		si, err := createServerInfo(pss, as.ProcNameParts, as.AlarmConditionWithWarningLevelChangeConditionMap)
		if err != nil {
			fmt.Println(errors.Wrap(err, 2).ErrorStack())
			panic(err)
		}

		err = common.SendPostWithJson(si, url)
		if err != nil {
			fmt.Println(errors.Wrap(err, 2).ErrorStack())
			panic(err)
		}

		fmt.Printf(".")
		i++
		if i%80 == 0 {
			fmt.Printf("\n")
		}
	}
}
