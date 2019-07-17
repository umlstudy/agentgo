package common

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ConvertStringToUtf8Bytes(text string) ([]byte, error) {
	var b bytes.Buffer
	wInUTF8 := transform.NewWriter(&b, unicode.UTF8.NewEncoder())
	defer wInUTF8.Close()
	_, err := wInUTF8.Write([]byte(text))
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func ConvertUtf8BytesToString(buf []byte) (string, error) {
	b := bytes.NewBuffer(buf)
	rInUTF8 := transform.NewReader(b, unicode.UTF8.NewDecoder())
	decBytes, err := ioutil.ReadAll(rInUTF8)
	if err != nil {
		return "", err
	}
	decS := string(decBytes)

	return decS, nil
}
