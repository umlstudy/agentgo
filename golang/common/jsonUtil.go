package common

import (
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
