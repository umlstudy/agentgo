package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// SendPostWithJSON is SendPostWithJSON
func SendPostWithJSON(any interface{}, url string) error {
	return sendJSON("POST", any, url)
}

// SendBytes is SendBytes
func SendBytes(method string, contentType string, byteArray []byte, url string) (retBody []byte, retStatusCode int, retErr error) {
	// JSON 바이트를 전송
	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteArray))
	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 500, errors.Wrap(err, "SendBytes")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, errors.Wrap(err, "SendBytes")
	}

	return body, resp.StatusCode, nil
}

func sendJSON(method string, any interface{}, url string) error {
	jsonBytes, err := ConvertObjectToJSONBytes(any)
	if err != nil {
		return errors.Wrap(err, "sendJson")
	}

	// JSON 바이트를 전송
	_, statusCode, err := SendBytes(method, "application/json", jsonBytes, url)
	if err != nil {
		return errors.Wrap(err, "sendJson")
	}
	if statusCode < 200 && statusCode >= 300 {
		return errors.New(fmt.Sprintf("response status is %d", statusCode))
	}

	return nil
}

func ReadBodyAsJson(res *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &result)
	return nil
}

// ResponseToJSON is ResponseToJSON
func ResponseToJSON(w http.ResponseWriter, object interface{}) error {
	js, err := ConvertObjectToJSONBytes(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
