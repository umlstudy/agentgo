package main

import (
	"fmt"
	"time"
)

func test9() {
	currTime := uint64(time.Now().Unix())
	fmt.Println(currTime)
	time.Sleep(5 * time.Second)
	currTime2 := uint64(time.Now().Unix())
	fmt.Println(currTime2 - currTime)
}
