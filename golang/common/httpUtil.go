package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func SendPostToJson(any interface{}, url string) error {
	return sendToJson(any, url, "POST")
}

func sendToJson(any interface{}, url string, method string) error {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		return err
	}

	// JSON 바이트를 문자열로 변경
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("response status is %d", resp.StatusCode))
	}

	return nil
}
