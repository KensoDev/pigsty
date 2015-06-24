package pigsty

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
)

type JobFlowRunner struct {
	JobFlowId *string
}

type RunJobFlowConfig struct {
	LogURI             string
	ClusterName        string
	KeyName            string
	InstanceCount      int64
	MasterInstanceType string
	SlaveInstanceType  string
	Region             string
	BootstrapActions   []BootstrapAction
}

func (config *RunJobFlowConfig) SetBootstrapActions(bootstrapActions []BootstrapAction) {
	config.BootstrapActions = bootstrapActions
}

func NewJobFlowConfig(config *RunJobFlowConfig) {

}

func (config *RunJobFlowConfig) RunJobFlow() (jobFlowRunner *JobFlowRunner) {
	runJobFlowInput := &emr.RunJobFlowInput{
		AMIVersion: aws.String("3.8.0"),
		Name:       aws.String(config.ClusterName),
		LogURI:     aws.String(config.LogURI),
		Instances: &emr.JobFlowInstancesConfig{
			EC2KeyName:         aws.String(config.KeyName),
			InstanceCount:      &config.InstanceCount,
			MasterInstanceType: aws.String(config.MasterInstanceType),
			SlaveInstanceType:  aws.String(config.SlaveInstanceType),
		},
	}

	for _, action := range config.BootstrapActions {
		panic(action)
	}

	svc := emr.New(&aws.Config{Region: config.Region})

	runJobFlowOutput, err := svc.RunJobFlow(runJobFlowInput)

	if err != nil {
		panic(err)
	}

	return &JobFlowRunner{JobFlowId: runJobFlowOutput.JobFlowID}
}
