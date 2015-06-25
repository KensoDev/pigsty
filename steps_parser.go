package pigsty

import (
	"encoding/json"
	"fmt"
)

type Step struct {
	Name            string
	ActionOnFailure string
	Args            []*string
	MainClass       string
	Jar             string
}

func NewStepsParser(jsonBlob []byte) (s []Step) {
	var steps []Step
	err := json.Unmarshal(jsonBlob, &steps)
	if err != nil {
		fmt.Println("error:", err)
	}
	return steps
}
