package main

import (
	"fmt"

	"os"

	"os/signal"

	"syscall"
)

func test11() {

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	fmt.Printf("pid is %d\n", os.Getpid())
	fmt.Printf("pid2 is %d\n", os.Getppid())
	fmt.Printf("pid2 is %d\n", os.Getuid())

	for {

		signal := <-signalChan

		switch signal {

		case syscall.SIGHUP:

			fmt.Printf("SIGHUP(%d)\n", signal)

		case syscall.SIGINT:

			fmt.Printf("SIGINT(%d)\n", signal)

		case syscall.SIGTERM:

			fmt.Printf("SIGTERM(%d)\n", signal)

		default:
			fmt.Printf("Unknown signal(%d)\n", signal)

		}

	}

}
