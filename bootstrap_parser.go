package pigsty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BootstrapAction struct {
	Name string
	Path string
	Args []*string
}

func NewBootstrapActions(jsonBlob []byte) (b []BootstrapAction) {
	var bootstrapActions []BootstrapAction
	err := json.Unmarshal(jsonBlob, &bootstrapActions)
	if err != nil {
		fmt.Println("error:", err)
	}
	return bootstrapActions
}

func ReadJsonFile(filename string) (content []byte) {
	if filename == "" {
		return []byte("")
	}
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Bootstrap config file unreadable, defaulting to blank file")
		content = []byte("")
	}

	return content
}
