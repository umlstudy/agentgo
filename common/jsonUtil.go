package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ReadJSON is ReadJSON
func ReadJSON(fileName string) (map[string]interface{}, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}

	var anyJSON map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &anyJSON)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}

	return anyJSON, nil
}

// ConvertObjectToJSONString is ConvertObjectToJSONString
func ConvertObjectToJSONString(any interface{}) (string, error) {
	jsonBytes, err := ConvertObjectToJSONBytes(any)
	if err != nil {
		return "", errors.Wrap(err, "ConvertObjectToJsonString")
	}
	var prettyJSONBuf bytes.Buffer
	err = json.Indent(&prettyJSONBuf, jsonBytes, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// ConvertObjectToJSONBytes is ConvertObjectToJSONBytes
func ConvertObjectToJSONBytes(any interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		return nil, errors.Wrap(err, "ConvertObjectToJsonBytes")
	}

	return jsonBytes, nil
}

func ConvertBytesToObject(buf []byte, result interface{}) error {
	// 동작안함
	json.Unmarshal(buf, &result)
	return nil
}
