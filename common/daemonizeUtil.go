package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type Runnable func()

func savePid(pid int, pidFile string, logger *log.Logger) error {

	file, err := os.Create(pidFile)
	if err != nil {
		return fmt.Errorf("Unable to create pid file : %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		return fmt.Errorf("Unable to create pid file : %v\n", err)
	}

	err = file.Sync() // flush to disk
	if err != nil {
		return err
	}

	return nil
}

func Daemonize(pidFile string, runnable Runnable, logger *log.Logger) {

	if strings.ToLower(os.Args[1]) == "main" {

		logger.Println("daemonize start")
		// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

		go func() {
			signalType := <-ch
			signal.Stop(ch)

			logger.Println("Exit command received. Exiting...")

			// this is a good place to flush everything to disk
			// before terminating.
			logger.Println("Received signal type : ", signalType)

			// remove PID file
			os.Remove(pidFile)

			os.Exit(0)

		}()

		runnable()

	} else if strings.ToLower(os.Args[1]) == "start" {

		// check if daemon already running.
		if _, err := os.Stat(pidFile); err == nil {
			logger.Printf("Already running or %s file exist.\n", pidFile)
			os.Exit(1)
		}

		args := []string{"main"}
		args = append(args, os.Args[2:]...)
		logger.Println(args)
		cmd := exec.Command(os.Args[0], args...)
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		logger.Println("Daemon process ID is : ", cmd.Process.Pid)
		err = savePid(cmd.Process.Pid, pidFile, logger)
		if err != nil {
			panic(err)
		}
		os.Exit(0)

	} else if strings.ToLower(os.Args[1]) == "stop" {
		// upon receiving the stop command
		// read the Process ID stored in PIDfile
		// kill the process using the Process ID
		// and exit. If Process ID does not exist, prompt error and quit

		if _, err := os.Stat(pidFile); err == nil {
			data, err := ioutil.ReadFile(pidFile)
			if err != nil {
				logger.Println("Not running")
				os.Exit(1)
			}
			ProcessID, err := strconv.Atoi(string(data))

			if err != nil {
				logger.Println("Unable to read and parse process id found in ", pidFile)
				os.Exit(1)
			}

			process, err := os.FindProcess(ProcessID)

			if err != nil {
				// remove PID file
				os.Remove(pidFile)

				logger.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			} else {
				// remove PID file
				os.Remove(pidFile)
			}

			logger.Printf("Killing process ID [%v] now.\n", ProcessID)
			// kill process and exit immediately
			err = process.Kill()

			if err != nil {
				logger.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
				os.Exit(1)
			} else {
				logger.Printf("Killed process ID [%v]\n", ProcessID)
				os.Exit(0)
			}
		} else {
			logger.Println("Not running.")
			os.Exit(1)
		}
	} else {
		logger.Printf("Unknown command : %v\n", os.Args[1])
		logger.Printf("Usage : %s [start|stop]\n", os.Args[0]) // return the program name back to %s
		os.Exit(1)
	}
}
