package pigsty

import (
	"encoding/json"
	"fmt"
)

type BootstrapAction struct {
	Path string
	Args []string
}

func NewBootstrapActionsParser(jsonBlob []byte) (b []BootstrapAction) {
	var bootstrapActions []BootstrapAction
	err := json.Unmarshal(jsonBlob, &bootstrapActions)
	if err != nil {
		fmt.Println("error:", err)
	}
	return bootstrapActions
}
