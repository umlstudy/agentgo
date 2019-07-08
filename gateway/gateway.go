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

var serverInfoMap = make(map[string]*ServerInfo)

func sayName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm a machine and my name is [miss.lee]"))
}

func responseServerInfos(w http.ResponseWriter, r *http.Request) {
	sis := []ServerInfo{}
	for _, value := range serverInfoMap {
		sis = append(sis, *value)
	}

	err := common.ResponseToJSON(w, sis)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

var lastAlarmTimes = map[string]uint64{}

func warningIfNeededTmp(as common.AbstractStatus, ID string) (needAlarm bool) {
	if as.WarningLevel == common.ERROR {
		if as.ResendAlarmLastSendAfter > 0 {
			judged := judgeAndSetLastResendBaseTime(ID, as.AlarmCondition)
			if judged {
				return true
			}
		}
	}
	return false
}

func judgeAndSetLastResendBaseTime(lastAlarmTimesKey string, ac common.AlarmCondition) bool {
	// 유효한 알람 발송 정보일 경우
	now := uint64(time.Now().Unix())
	lastResendBaseTime := (lastAlarmTimes[lastAlarmTimesKey] + ac.SendAlarmOccuredAfter + ac.ResendAlarmLastSendAfter + (60 * 5))
	if lastResendBaseTime < now {
		lastAlarmTimes[lastAlarmTimesKey] = now
		return true
	}
	return false
}

func warningIfNeeded(si common.ServerInfo) {
	var alarmMessages bytes.Buffer

	// 1.1
	needAlarm := !si.IsRunning
	if !si.IsRunning {
		judged := judgeAndSetLastResendBaseTime("host", si.AlarmCondition)
		if judged {
			alarmMessages.WriteString("system died.")
		}
	}

	// 1.2
	for _, rs := range si.ResourceStatuses {
		as := rs.AbstractStatus
		currNeedAlarm := warningIfNeededTmp(as, fmt.Sprintf("%s-%s", si.ID, rs.ID))
		needAlarm = needAlarm || currNeedAlarm
		if as.WarningLevel == common.WARNING {
			alarmMessages.WriteString(fmt.Sprintf(", resource warning (%s)", rs.Name))
		}
		if as.WarningLevel == common.ERROR {
			alarmMessages.WriteString(fmt.Sprintf(", resource error (%s)", rs.Name))
		}
	}

	// 1.3
	for _, ps := range si.ProcessStatuses {
		as := ps.AbstractStatus
		currNeedAlarm := warningIfNeededTmp(as, fmt.Sprintf("%s-%s", si.ID, ps.ID))
		needAlarm = needAlarm || currNeedAlarm
		if as.WarningLevel == common.WARNING {
			alarmMessages.WriteString(fmt.Sprintf(", resource warning (%s)", ps.Name))
		}
		if as.WarningLevel == common.ERROR {
			alarmMessages.WriteString(fmt.Sprintf(", resource error (%s)", ps.Name))
		}
	}

	// 2
	if needAlarm {
		// 알람 실행
		fmt.Printf("--\n--\n--\n-- ALARM - %s\nalarm %s\n", time.Now(), alarmMessages.String())
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

func getColorString2(si common.ServerInfo) string {
	if si.IsRunning {
		return "green"
	}
	return "red"
}
func displayStatusDetail(si *common.ServerInfo) {

	color := getColorString2(*si)
	siNameMsg := si.Name
	if !si.IsRunning {
		siNameMsg = fmt.Sprintf("%s-Died", siNameMsg)
	}
	colorstring.Fprintf(ansi.NewAnsiStdout(), fmt.Sprintf("[%s]* %s\n", color, siNameMsg))
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

// AnsiLine is AnsiLine
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

var lastServerInfoRecvTimeMap = map[string]uint64{}

var enableDisplay = false

// var serverJudgeDiedTime = uint64(60*5) // 5분
var serverJudgeDiedTime = uint64(20) // 20초

func runLoop(runLoopQuitChan <-chan bool) {
	notFinished := true
	for notFinished {
		select {
		case <-runLoopQuitChan:
			notFinished = false
			break
		default:
			ansiLine.reset()
			// 1.서버 상태 변경
			if len(serverInfoMap) > 0 {
				for _, si := range serverInfoMap {
					lastServerInfoRecvTime := lastServerInfoRecvTimeMap[si.ID]
					currTime := uint64(time.Now().Unix())
					// 10분 이상 serverInfo 가 갱신이 되지 않았다면
					// 서버와의 연결이 종료된 것으로 간주함
					if (currTime - lastServerInfoRecvTime) > serverJudgeDiedTime {
						si.IsRunning = false
					}
				}
			}

			// 2.DISPLAY
			if enableDisplay {
				if len(serverInfoMap) > 0 {
					for _, si := range serverInfoMap {
						displayStatusDetail(si)
					}
				} else {
					colorstring.Fprintf(ansi.NewAnsiStdout(), "[red][bold]empty serverInfos")
					ansi.CursorHorizontalAbsolute(0)
				}
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

	lastServerInfoRecvTimeMap[si.ID] = uint64(time.Now().Unix())
	serverInfoMap[si.ID] = &si

	err = common.ResponseToJSON(rw, "OK")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	var port = flag.Int("port", common.DefaultServerPort, "ServerMonitory Gateway's port no")
	var enableDisplayPtr = flag.Bool("enableDisplay", false, "Enable display for ServerMonitory Gateway")
	flag.Parse()

	enableDisplay = *enableDisplayPtr

	fmt.Printf("> Using port is %v\n", *port)
	fmt.Printf("> EnableDisplay is %v\n", enableDisplay)

	webServerStart(*port)
}

func webServerStart(port int) {

	runLoopQuitChan := make(chan bool)
	defer close(runLoopQuitChan)

	mux := http.NewServeMux()

	mux.HandleFunc("/", sayName)                               // 살아있는지 테스트용
	mux.HandleFunc("/getServerInfos", responseServerInfos)     // 모니터UI로부터의 응답
	mux.HandleFunc("/recvServerInfo", recvServerInfoFromAgent) // 에이전트로부터의 자료 수신

	t := time.Now()
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> ServerMonitory Gateway Start at %s\n", t.Format("2006-01-02 15:04:05"))
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> Waiting for agent or front end UI...\n")

	ansi.CursorHide()

	go runLoop(runLoopQuitChan)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)

	// 종료 시그널 보냄
	runLoopQuitChan <- true
	if err != nil {
		panic(err)
	}
	colorstring.Fprintf(ansi.NewAnsiStdout(), "[blue][bold]> ServerMonitory Gateway Stop at %s\n", t.Format("2006-01-02 15:04:05"))
	ansi.CursorShow()
}
