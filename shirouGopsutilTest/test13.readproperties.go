package main

import (
	"fmt"
	"github.com/umlstudy/serverMonitor/common"
)

func test13() {
	props, err := common.ReadPropertiesFile("test13.properties")
    if err != nil {
		panic(err)
    }

	for k,v := range props {
		fmt.Printf("%s,%s\n",k,v)
	}
}
