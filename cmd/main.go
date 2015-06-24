package main

import (
	"fmt"
	"github.com/kensodev/pigsty"
)

func main() {
	content := pigsty.ReadJsonFile("bootstrap-actions-sample.json")
	bootstrapActions := pigsty.NewBootstrapActions(content)

	fmt.Printf("You have %v Bootstrap actions", len(bootstrapActions))
}

func runCluster() {
	config := &pigsty.RunJobFlowConfig{
		LogURI:             "s3://metasearch-impressions",
		ClusterName:        "metasearch-impressions",
		KeyName:            "gogobot-production",
		InstanceCount:      3,
		MasterInstanceType: "m3.xlarge",
		SlaveInstanceType:  "m3.xlarge",
	}

	runner := pigsty.NewJobFlowRunner(config)

	fmt.Println("Running: " + *runner.JobFlowId)
}
