package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJson(fileName string) (map[string]interface{}, error) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var anyJson map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &anyJson)
	if err != nil {
		return nil, err
	}

	return anyJson, nil
}
