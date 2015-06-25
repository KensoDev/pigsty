package main

import (
	"fmt"
	"github.com/kensodev/pigsty"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	logURI              = kingpin.Flag("log-uri", "Logs location for pig/hadoop. Usually this will point to an S3 bucket").String()
	clusterName         = kingpin.Flag("cluster-name", "What should be the cluster name").String()
	keyName             = kingpin.Flag("key-name", "Which key-pair should the instances use").String()
	instanceCount       = kingpin.Flag("instance-count", "How many instances should be boot up").Int()
	masterInstanceType  = kingpin.Flag("master-instance-type", "What should be the instance type for the master").String()
	slaveInstanceType   = kingpin.Flag("slave-instance-type", "What should be the instance type for the slave").String()
	bootstrapConfigFile = kingpin.Flag("bootstrap-configuration-file", "Configuration for bootstrap steps file location").String()
	stepsConfigFile     = kingpin.Flag("steps-configuration-file", "Configuration for steps steps file location").String()
	region              = kingpin.Flag("region", "EC2 Region").String()
)

func main() {
	kingpin.Parse()

	bootstrapReader := pigsty.JsonReader{FileName: *bootstrapConfigFile}
	bootstrapActions := pigsty.NewBootstrapActions(bootstrapReader.ReadFile())

	stepsReader := pigsty.JsonReader{FileName: *stepsConfigFile}
	steps := pigsty.NewStepsParser(stepsReader.ReadFile())

	config := &pigsty.RunJobFlowConfig{
		LogURI:             *logURI,
		ClusterName:        *clusterName,
		KeyName:            *keyName,
		InstanceCount:      *instanceCount,
		MasterInstanceType: *masterInstanceType,
		SlaveInstanceType:  *slaveInstanceType,
		Region:             *region,
		BootstrapActions:   bootstrapActions,
		Steps:              steps,
	}

	runCluster(config)
}

func runCluster(config *pigsty.RunJobFlowConfig) {
	runner := config.RunJobFlow()

	fmt.Println("Running: " + *runner.JobFlowId)
}
