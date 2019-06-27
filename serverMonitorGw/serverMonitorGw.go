package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResourceStatus struct {
	Id    string `json:"id"`
	Min   uint32 `json:"min"`
	Max   uint32 `json:"max"`
	Name  string `json:"name"`
	Value uint32 `json:"value"`
}
type ServerInfo struct {
	Id               string           `json:"id"`
	Name             string           `json:"name"`
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
}

var serverInfoMap = map[string]ServerInfo{}

func BeforeStart() {
	serverInfoMap["mysvr"] = ServerInfo{"mysvr", "mysvr", []ResourceStatus{
		ResourceStatus{"cpu", 1, 100, "cpu", 1},
		ResourceStatus{"mem", 1, 100, "memory", 1},
		ResourceStatus{"di1", 1, 100, "disk1", 1},
		ResourceStatus{"di2", 1, 100, "disk2", 1},
	}}
}

func SayName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm a machine and my name is [miss.lee]"))
}

func SendServerInfos(w http.ResponseWriter, r *http.Request) {
	sis := []ServerInfo{}
	for _, value := range serverInfoMap {
		sis = append(sis, value)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	Send(w, r, sis)
}

func Send(w http.ResponseWriter, r *http.Request, x interface{}) {
	js, err := json.Marshal(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var recvCnt = 0

func RecvServerInfo(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var si ServerInfo
	err := decoder.Decode(&si)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	serverInfoMap[si.Id] = si

	Send(rw, req, "OK")

	recvCnt++
	fmt.Printf(".")
	if recvCnt%80 == 0 {
		fmt.Printf("\n")
	}
}

func main() {
	BeforeStart()
	WebServerStart()
}

func WebServerStart() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", SayName)
	mux.HandleFunc("/getServerInfos", SendServerInfos)
	mux.HandleFunc("/recvServerInfo", RecvServerInfo)

	t := time.Now()
	fmt.Printf("> start %s\n", t.Format("2006-01-02 15:04:05"))
	http.ListenAndServe(":8080", mux)
	fmt.Printf("> end %s\n", t.Format("2006-01-02 15:04:05"))
}
