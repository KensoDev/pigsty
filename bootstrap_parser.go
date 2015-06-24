package pigsty

import (
	"encoding/json"
	"fmt"
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
