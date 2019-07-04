package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"sejong.asia/serverMonitor/common"
)

// ServerInfo is ServerInfo
type ServerInfo = common.ServerInfo

var serverInfoMap = map[string]ServerInfo{}

func sayName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm a machine and my name is [miss.lee]"))
}

func responseServerInfos(w http.ResponseWriter, r *http.Request) {
	sis := []ServerInfo{}
	for _, value := range serverInfoMap {
		sis = append(sis, value)
	}

	err := common.ResponseToJson(w, sis)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

var recvCnt = 0

var lastAlarmTimes = map[string]uint64{}

func warningIfNeeded(as common.AbstractStatus, id string) (bool, bool, string) {
	if as.WarningLevel == common.ERROR || as.WarningLevel == common.WARNING {
		if as.ResendAlarmLastSendAfter > 0 {
			// 유효한 알람 발송 정보일 경우
			now := uint64(time.Now().Unix())
			lastResendBaseTime := (lastAlarmTimes[id] + as.SendAlarmOccuredAfter + as.ResendAlarmLastSendAfter + (60 * 5))

			if lastResendBaseTime < now {
				lastAlarmTimes[id] = now
				return true, true, ""
			}
		}
	}
	return true, true, ""
}

func recvServerInfoFromAgent(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var si ServerInfo
	err := decoder.Decode(&si)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		fmt.Printf("%v\n", err)
		return
	}

	warningIfNeeded(si.AbstractStatus, si.Id)

	if recvCnt%20 == 0 {
		jsonStr, err := common.ConvertObjectToJsonString(si)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", jsonStr)
	}

	serverInfoMap[si.Id] = si

	err = common.ResponseToJson(rw, "OK")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	recvCnt++

	fmt.Printf(".")
	if recvCnt%80 == 0 {
		fmt.Printf("\n")
	}
}

func main() {
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()
	webServerStart(*port)
}

func webServerStart(port int) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", sayName)                               // 살아있는지 테스트용
	mux.HandleFunc("/getServerInfos", responseServerInfos)     // 모니터UI로 자료 전송
	mux.HandleFunc("/recvServerInfo", recvServerInfoFromAgent) // 에이전트로부터 자료 수신

	t := time.Now()
	fmt.Printf("> ServerMonitory Gateway Start at %s\n", t.Format("2006-01-02 15:04:05"))
	fmt.Printf(fmt.Sprintf("> Waiting for agent or front end UI... (port:%d)\n", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		panic(err)
	}
	fmt.Printf("> ServerMonitory Gateway Stop at %s\n", t.Format("2006-01-02 15:04:05"))
}
