package main

import (
	"flag"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var pidFile = "/tmp/daemonize2.pid"
var logger *log.Logger

func savePID(pid int) {

	file, err := os.Create(pidFile)
	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk

}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World"

	w.Write([]byte(html))
}

func initLog() {
}

func main10() {
	fpLog, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	logger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	f1 := flag.NewFlagSet("f1", flag.ExitOnError)
	silent := f1.Bool("silent", false, "")
	f2 := flag.NewFlagSet("f2", flag.ContinueOnError)
	loud := f2.Bool("loud", false, "")

	subCmd := ""
	if len(os.Args) >= 2 {
		subCmd = os.Args[1]
	}
	switch subCmd {
	case "main":
	case "start":
		if err := f1.Parse(os.Args[2:]); err == nil {
			fmt.Println("apply", *silent)
		}
	case "stop":
		if err := f2.Parse(os.Args[2:]); err == nil {
			fmt.Println("reset", *loud)
		}
	default:
		fmt.Printf("UsageX : %s [start|stop] \n ", os.Args[0])
		os.Exit(0)
	}

	daemonize(pidFile, mainExec)
}

type runnable func()

func mainExec() {
	fmt.Println("START")
	logger.Println("mainExec start")
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHelloWorld)
	log.Fatalln(http.ListenAndServe(":8080", mux))
	logger.Println("end")
}

func daemonize(pidFile string, runFunc runnable) {

	if strings.ToLower(os.Args[1]) == "main" {
		initLog()

		logger.Println("main STAR")
		// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

		go func() {
			signalType := <-ch
			signal.Stop(ch)
			logger.Println("Exit command received. Exiting...")

			// this is a good place to flush everything to disk
			// before terminating.
			fmt.Println("Received signal type : ", signalType)

			// remove PID file
			os.Remove(pidFile)

			os.Exit(0)

		}()

		runFunc()

	} else if strings.ToLower(os.Args[1]) == "start" {

		// check if daemon already running.
		if _, err := os.Stat(pidFile); err == nil {
			fmt.Printf("Already running or %s file exist.\n", pidFile)
			os.Exit(1)
		}

		args := []string{"main"}
		args = append(args, os.Args[2:]...)
		fmt.Println(args)
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
		savePID(cmd.Process.Pid)
		os.Exit(0)

	} else if strings.ToLower(os.Args[1]) == "stop" {
		// upon receiving the stop command
		// read the Process ID stored in PIDfile
		// kill the process using the Process ID
		// and exit. If Process ID does not exist, prompt error and quit

		if _, err := os.Stat(pidFile); err == nil {
			data, err := ioutil.ReadFile(pidFile)
			if err != nil {
				fmt.Println("Not running")
				os.Exit(1)
			}
			ProcessID, err := strconv.Atoi(string(data))

			if err != nil {
				fmt.Println("Unable to read and parse process id found in ", pidFile)
				os.Exit(1)
			}

			process, err := os.FindProcess(ProcessID)

			if err != nil {
				// remove PID file
				os.Remove(pidFile)

				fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			} else {
				// remove PID file
				os.Remove(pidFile)
			}

			fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
			// kill process and exit immediately
			err = process.Kill()

			if err != nil {
				fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			} else {
				fmt.Printf("Killed process ID [%v]\n", ProcessID)
				os.Exit(0)
			}
		} else {
			fmt.Println("Not running.")
			os.Exit(1)
		}
	} else {
		fmt.Printf("Unknown command : %v\n", os.Args[1])
		fmt.Printf("Usage : %s [start|stop]\n", os.Args[0]) // return the program name back to %s
		os.Exit(1)
	}
}
