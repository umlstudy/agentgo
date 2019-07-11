package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	HostId                                           string                                                          `json:"hostId"`
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

func createServerInfoAndRetry(hostId string, pss []common.ProcessStatus, procNameParts []string, alarmConditionWithWarningLevelChangeConditionMap map[string]common.AlarmConditionWithWarningLevelChangeCondition) (*common.ServerInfo, error) {
	si, err := createServerInfo(hostId, pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
	retryCnt := 0
	for err != nil && retryCnt < 3 {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		time.Sleep(5 * time.Second)
		si, err = createServerInfo(hostId, pss, procNameParts, alarmConditionWithWarningLevelChangeConditionMap)
	}
	return si, err
}

var logger *log.Logger

func main() {

	gatewayHost := flag.String("host", "localhost", "ServerMonitory Gateway's host name or ip to gateway")
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	enableConsoleLogPtr := flag.Bool("enableConsoleLog", false, "Enable console log for ServerMonitory Gateway")
	flag.Parse()

	fmt.Printf("> Using port is %v\n", *port)
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[green][bold]> Target host is %v\n", *gatewayHost)
	fmt.Printf("> EnableConsoleLog is %v\n", *enableConsoleLogPtr)
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

	const logParam = log.Ldate | log.Ltime | log.Lshortfile
	if !*enableConsoleLogPtr {
		fpLog, err := os.OpenFile("agent.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fpLog.Close()

		logger = log.New(fpLog, "", logParam)
		runMain(gatewayHost, port, pss, as)
	} else {
		logger = log.New(os.Stdout, "", logParam)
		runMain(gatewayHost, port, pss, as)
	}
}

func runMain(gatewayHost *string, port *int, pss []common.ProcessStatus, as *AgentSettings) {
	url := fmt.Sprintf(urlFormat, *gatewayHost, *port)
	logger.Println("> Agent started...")
	for true {

		time.Sleep(5 * time.Second)

		si, err := createServerInfoAndRetry(as.HostId, pss, as.ProcNameParts, as.AlarmConditionWithWarningLevelChangeConditionMap)
		if err != nil {
			logger.Println(errors.Wrap(err, 2).ErrorStack())
			panic(err)
		}

		err = common.SendPostWithJSON(si, url)
		if err != nil {
			logger.Println(errors.Wrap(err, 2).ErrorStack())
			logger.Println("Server is Not Ready... Wait 10 Second and retry...")
			time.Sleep(10 * time.Second)
			continue
		}
	}
}
