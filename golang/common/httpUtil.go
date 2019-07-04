package common

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func SendPostWithJson(any interface{}, url string) error {
	return sendJson("POST", any, url)
}

func SendBytes(method string, byteArray []byte, url string) (*http.Response, error) {
	// JSON 바이트를 전송
	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteArray))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "SendBytes")
	}

	defer resp.Body.Close()

	return resp, nil
}

func sendJson(method string, any interface{}, url string) error {
	jsonBytes, err := ConvertObjectToJsonBytes(any)
	if err != nil {
		return errors.Wrap(err, "sendJson")
	}

	// JSON 바이트를 전송
	resp, err := SendBytes(method, jsonBytes, url)
	if err != nil {
		return errors.Wrap(err, "sendJson")
	}
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("response status is %d", resp.StatusCode))
	}

	return nil
}
