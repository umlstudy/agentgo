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

// // ResourceStatus is ResourceStatus
// type ResourceStatus = common.ResourceStatus

// // ProcessStatus is ProcessStatus
// type ProcessStatus = common.ProcessStatus

// func beforeStart() {
// 	serverInfoMap["mysvr"] = ServerInfo{"mysvr", "mysvr",
// 		[]ResourceStatus{
// 			ResourceStatus{"cpu", 1, 100, "cpu", 1},
// 			ResourceStatus{"mem", 1, 100, "memory", 1},
// 			ResourceStatus{"di1", 1, 100, "disk1", 1},
// 			ResourceStatus{"di2", 1, 100, "disk2", 1},
// 		},
// 		[]ProcessStatus{
// 			ProcessStatus{"tomcat", "tomcat", "tomcat", 1111},
// 		}}
// }

var serverInfoMap = map[string]ServerInfo{}

func sayName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm a machine and my name is [miss.lee]"))
}

func sendServerInfos(w http.ResponseWriter, r *http.Request) {
	sis := []ServerInfo{}
	for _, value := range serverInfoMap {
		sis = append(sis, value)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	send(w, r, sis)
}

func send(w http.ResponseWriter, r *http.Request, x interface{}) {
	js, err := json.Marshal(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var recvCnt = 0

func recvServerInfo(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var si ServerInfo
	err := decoder.Decode(&si)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	serverInfoMap[si.Id] = si

	send(rw, req, "OK")

	recvCnt++
	fmt.Printf(".")
	if recvCnt%80 == 0 {
		fmt.Printf("\n")
	}
}

func main() {
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()

	// beforeStart()
	webServerStart(*port)
}

func webServerStart(port int) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", sayName)                       // 살아있는지 테스트용
	mux.HandleFunc("/getServerInfos", sendServerInfos) // 모니터UI로 자료 전송
	mux.HandleFunc("/recvServerInfo", recvServerInfo)  // 에이전트로부터 자료 수신

	t := time.Now()
	fmt.Printf("> ServerMonitory Gateway Start at %s\n", t.Format("2006-01-02 15:04:05"))
	fmt.Printf(fmt.Sprintf("> Waiting for agent or front end UI... (port:%d)\n", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		panic(err)
	}
	fmt.Printf("> ServerMonitory Gateway Stop at %s\n", t.Format("2006-01-02 15:04:05"))
}
