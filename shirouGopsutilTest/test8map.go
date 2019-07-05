package main

import (
	"fmt"
)

func test8() {
	mymap := map[string]string{}
	mymap["aaa"] = "bbb"
	mymap["aaa1"] = "bbb1"
	mymap["aaa4"] = "bbb4"
	mymap["aaa2"] = "bbb2"
	mymap["aaa3"] = "bbb3"

	for k := range mymap {
		if k == "aaa4" {
			delete(mymap, k)
		}
	}

	for _, v := range mymap {
		fmt.Println(v)
	}
}
