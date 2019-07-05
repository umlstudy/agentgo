package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
	"github.com/umlstudy/serverMonitor/common"
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

// var recvCnt = 0

var lastAlarmTimes = map[string]uint64{}

func warningIfNeeded_(as common.AbstractStatus, id string) (needAlarm bool) {
	if as.WarningLevel == common.ERROR || as.WarningLevel == common.WARNING {
		if as.ResendAlarmLastSendAfter > 0 {
			// 유효한 알람 발송 정보일 경우
			now := uint64(time.Now().Unix())
			lastResendBaseTime := (lastAlarmTimes[id] + as.SendAlarmOccuredAfter + as.ResendAlarmLastSendAfter + (60 * 5))

			if lastResendBaseTime < now {
				lastAlarmTimes[id] = now
				return true
			}
		}
	}
	return false
}

func warningIfNeeded(si common.ServerInfo) {
	var buffer bytes.Buffer

	// 1.
	as := si.AbstractStatus
	needAlarm := warningIfNeeded_(as, si.Id)
	if as.WarningLevel == common.WARNING || as.WarningLevel == common.ERROR {
		buffer.WriteString("\nsystem error occured")
	}

	for _, rs := range si.ResourceStatuses {
		as := rs.AbstractStatus
		needAlarm = needAlarm || warningIfNeeded_(as, fmt.Sprintf("%s-%s", si.Id, rs.Id))
		if as.WarningLevel == common.WARNING || as.WarningLevel == common.ERROR {
			buffer.WriteString(fmt.Sprintf("\nresource error (%s)", rs.Name))
		}
	}

	for _, ps := range si.ProcessStatuses {
		as := ps.AbstractStatus
		needAlarm = needAlarm || warningIfNeeded_(as, fmt.Sprintf("%s-%s", si.Id, ps.Id))
		if as.WarningLevel == common.WARNING || as.WarningLevel == common.ERROR {
			buffer.WriteString(fmt.Sprintf("\nresource error (%s)", ps.Name))
		}
	}

	if needAlarm {
		// 알람 실행
		fmt.Printf("%s\n", buffer.String())
	}
}

func getColorString(wl common.WarningLevel) string {
	if wl == common.WARNING {
		return "yellow"
	} else if wl == common.ERROR {
		return "red"
	} else {
		return "green"
	}
}

func displayStatus_(si common.ServerInfo) {

	as := si.AbstractStatus
	color := getColorString(as.WarningLevel)
	colorstring.Fprintf(ansi.NewAnsiStdout(), fmt.Sprintf("[%s]* %s\n", color, si.Name))
	ansiLine.inc()

	for idx, rs := range si.ResourceStatuses {
		if idx > 0 {
			fmt.Printf(", ")
		}
		as := rs.AbstractStatus
		color = getColorString(as.WarningLevel)
		formatString := fmt.Sprintf("[%s]%s(%d)", color, rs.Name, rs.Value)
		colorstring.Fprintf(ansi.NewAnsiStdout(), formatString)
	}
	fmt.Printf("\n")
	ansiLine.inc()

	for idx, ps := range si.ProcessStatuses {
		if idx > 0 {
			fmt.Printf(", ")
		}
		as := ps.AbstractStatus
		color = getColorString(as.WarningLevel)
		formatString := fmt.Sprintf("[%s]%s", color, ps.Name)
		colorstring.Fprintf(ansi.NewAnsiStdout(), formatString)
	}
	fmt.Printf("\n")
	ansiLine.inc()
}

type AnsiLine struct {
	currentLine uint32
	lastLine    uint32
}

func (al *AnsiLine) inc() {
	al.currentLine++
	if al.currentLine > al.lastLine {
		al.lastLine = al.currentLine
	}
}

func (al *AnsiLine) reset() {
	ansi.CursorUp(int(al.currentLine))
	al.currentLine = 0
}

var ansiLine = AnsiLine{}

func displayStatus(quitRecvChan <-chan bool) {
	notFinished := true
	for notFinished {
		select {
		case <-quitRecvChan:
			notFinished = false
			break
		default:
			ansiLine.reset()
			if len(serverInfoMap) > 0 {
				for _, si := range serverInfoMap {
					displayStatus_(si)
				}
			} else {
				colorstring.Fprintf(ansi.NewAnsiStdout(), "[red][bold]empty serverInfos")
				ansi.CursorHorizontalAbsolute(0)
			}
			time.Sleep(2 * time.Second)
			break
		}
	}
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

	warningIfNeeded(si)

	serverInfoMap[si.Id] = si

	err = common.ResponseToJson(rw, "OK")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	port := flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	flag.Parse()
	webServerStart(*port)
}

func webServerStart(port int) {

	quitRecvChan := make(chan bool)
	defer close(quitRecvChan)

	mux := http.NewServeMux()

	mux.HandleFunc("/", sayName)                               // 살아있는지 테스트용
	mux.HandleFunc("/getServerInfos", responseServerInfos)     // 모니터UI로 자료 전송
	mux.HandleFunc("/recvServerInfo", recvServerInfoFromAgent) // 에이전트로부터 자료 수신

	t := time.Now()
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> ServerMonitory Gateway Start at %s\n", t.Format("2006-01-02 15:04:05"))
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> Waiting for agent or front end UI... (port:%d)\n", port)

	ansi.CursorHide()

	go displayStatus(quitRecvChan)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	// 종료 시그널 보냄
	quitRecvChan <- true
	if err != nil {
		panic(err)
	}
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> ServerMonitory Gateway Stop at %s\n", t.Format("2006-01-02 15:04:05"))
	ansi.CursorShow()
}
