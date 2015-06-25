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
	InstanceCount      int
	MasterInstanceType string
	SlaveInstanceType  string
	Region             string
	BootstrapActions   []BootstrapAction
	Steps              []Step
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
			InstanceCount:      aws.Long(int64(config.InstanceCount)),
			MasterInstanceType: aws.String(config.MasterInstanceType),
			SlaveInstanceType:  aws.String(config.SlaveInstanceType),
		},
	}
	runJobFlowInput.BootstrapActions = []*emr.BootstrapActionConfig{}

	for _, bootstrapAction := range config.BootstrapActions {
		action := &emr.BootstrapActionConfig{
			Name: aws.String(bootstrapAction.Name),
			ScriptBootstrapAction: &emr.ScriptBootstrapActionConfig{
				Path: aws.String(bootstrapAction.Path),
				Args: bootstrapAction.Args,
			},
		}
		runJobFlowInput.BootstrapActions = append(runJobFlowInput.BootstrapActions, action)
	}

	runJobFlowInput.Steps = []*emr.StepConfig{}

	for _, step := range config.Steps {
		stepConfig := &emr.StepConfig{
			Name:            aws.String(step.Name),
			ActionOnFailure: aws.String(step.ActionOnFailure),
			HadoopJARStep: &emr.HadoopJARStepConfig{
				JAR:       aws.String(step.Jar),
				MainClass: aws.String(step.MainClass),
				Args:      step.Args,
			},
		}
		runJobFlowInput.Steps = append(runJobFlowInput.Steps, stepConfig)
	}

	svc := emr.New(&aws.Config{Region: config.Region})

	runJobFlowOutput, err := svc.RunJobFlow(runJobFlowInput)

	if err != nil {
		panic(err)
	}

	return &JobFlowRunner{JobFlowId: runJobFlowOutput.JobFlowID}
}
