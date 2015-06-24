package pigsty

import (
	"fmt"
	"io/ioutil"
)

type JsonReader struct {
	FileName string
}

func (reader *JsonReader) ReadFile() (content []byte) {
	if reader.FileName == "" {
		return []byte("")
	}
	content, err := ioutil.ReadFile(reader.FileName)

	if err != nil {
		fmt.Println("Bootstrap config file unreadable, defaulting to blank file")
		content = []byte("")
	}

	return content
}
