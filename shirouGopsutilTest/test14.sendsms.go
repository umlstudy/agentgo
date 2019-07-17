package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/umlstudy/serverMonitor/common"
)

type SmsResponse struct {
	Result string `json:"result"`
	Type   string `json:"type"`
	MsgId  string `json:"msgid"`
	OkCnt  uint32 `json:"ok_cnt"`
}

func test14() {
	props, err := common.ReadPropertiesFile("test13.properties")
	if err != nil {
		panic(err)
	}

	userid := props["userid"]
	callback := props["callback"]
	phone := props["phone"]
	msg := props["msg"]
	bodyString := fmt.Sprintf("userid=%s&callback=%s&phone=%s&msg=%s",
		userid, callback, phone, msg)
	//byteArray, err := common.ConvertStringToUtf8Bytes(bodyString)
	byteArray := []byte(bodyString)
	fmt.Println(bodyString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", byteArray)
	url := props["url"]
	fmt.Println(url)
	body, _, err := common.SendBytes("POST", "application/x-www-form-urlencoded", byteArray, url)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("d:/abc.txt", body, 0644)
	if err != nil {
		panic(err)
	}

	rstr, err := common.ConvertUtf8BytesToString(body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("0 %v\n", body)
	fmt.Printf("3 %v\n", rstr)

	var data SmsResponse
	fmt.Printf("1 %v\n", data)
	json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("2 %v\n", data)

	rstr, err = common.ConvertUtf8BytesToString(byteArray)
	if err != nil {
		panic(err)
	}
	fmt.Printf("4 %v\n", rstr)
}

// func fileW() {

//     // To start, here's how to dump a string (or just
//     // bytes) into a file.
//     d1 := []byte("hello\ngo\n")
//     err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
//     check(err)

//     // For more granular writes, open a file for writing.
//     f, err := os.Create("/tmp/dat2")
//     check(err)

//     // It's idiomatic to defer a `Close` immediately
//     // after opening a file.
//     defer f.Close()

//     // You can `Write` byte slices as you'd expect.
//     d2 := []byte{115, 111, 109, 101, 10}
//     n2, err := f.Write(d2)
//     check(err)
//     fmt.Printf("wrote %d bytes\n", n2)

//     // A `WriteString` is also available.
//     n3, err := f.WriteString("writes\n")
//     fmt.Printf("wrote %d bytes\n", n3)

//     // Issue a `Sync` to flush writes to stable storage.
//     f.Sync()

//     // `bufio` provides buffered writers in addition
//     // to the buffered readers we saw earlier.
//     w := bufio.NewWriter(f)
//     n4, err := w.WriteString("buffered\n")
//     fmt.Printf("wrote %d bytes\n", n4)

//     // Use `Flush` to ensure all buffered operations have
//     // been applied to the underlying writer.
//     w.Flush()

// }
