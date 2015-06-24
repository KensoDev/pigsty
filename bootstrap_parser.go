package pigsty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BootstrapAction struct {
	Path string
	Args []string
}

func (f *Foo) SetName(name string) {
	f.name = name
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

	content, err := ioutil.ReadFile("bootstrap-actions-sample.json")

	if err != nil {
		panic("Bootstrap config file unreadable")
	}

	return content
}
