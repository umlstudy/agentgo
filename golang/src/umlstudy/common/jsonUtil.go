package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func ReadJson(fileName string) (map[string]interface{}, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}

	var anyJson map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &anyJson)
	if err != nil {
		return nil, errors.Wrap(err, "readJson")
	}

	return anyJson, nil
}

func ConvertObjectToJsonString(any interface{}) (string, error) {
	jsonBytes, err := ConvertObjectToJsonBytes(any)
	if err != nil {
		return "", errors.Wrap(err, "ConvertObjectToJsonString")
	}
	var prettyJsonBuf bytes.Buffer
	err = json.Indent(&prettyJsonBuf, jsonBytes, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func ConvertObjectToJsonBytes(any interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		return nil, errors.Wrap(err, "ConvertObjectToJsonBytes")
	}

	return jsonBytes, nil
}
