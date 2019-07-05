package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-errors/errors"
	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
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

func createServerInfoAndRetry(pss []common.ProcessStatus, procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) (*common.ServerInfo, error) {
	si, err := createServerInfo(pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
	retryCnt := 0
	for err != nil && retryCnt < 3 {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		time.Sleep(5 * time.Second)
		si, err = createServerInfo(pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
	}
	return si, err
}

func main() {

	host := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()

	fmt.Printf("> Using port is %v\n", *port)
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[green][bold]> Target host is %v\n", *host)

	fmt.Printf("> Reading agentSettings.json...\n")
	as, err := readJSON("agentSettings.json")
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		panic(err)
	}
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[green][bold]> PASS Reading agentSettings.json...\n")

	fmt.Printf("> Scanning process ids...\n")
	pss, err := findMatchedPids(as.ProcNameParts, as.AlarmConditionWithWarningLevelChangeConditionMap)
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		panic(err)
	}
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[green][bold]> PASS scanning process ids...\n")

	url := fmt.Sprintf(urlFormat, *host, *port)
	i := 0
	for true {

		time.Sleep(5 * time.Second)

		si, err := createServerInfoAndRetry(pss, as.ProcNameParts, as.AlarmConditionWithWarningLevelChangeConditionMap)
		if err != nil {
			fmt.Println(errors.Wrap(err, 2).ErrorStack())
			panic(err)
		}

		err = common.SendPostWithJSON(si, url)
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
