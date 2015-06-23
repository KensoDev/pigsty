package pigsty

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"time"
)

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := emr.New(&aws.Config{Region: "us-west-2"})

	clusterId := aws.String("j-A0KZX865NTUD")

	resp, err := svc.DescribeCluster(&emr.DescribeClusterInput{ClusterID: clusterId})
	if err != nil {
		panic(err)
	}

	clusterState := *resp.Cluster.Status.State

	fmt.Println("Cluster Status id: " + clusterState)

	for {
		resp, err = svc.DescribeCluster(&emr.DescribeClusterInput{ClusterID: clusterId})
		if err != nil {
			panic(err)
		}
		clusterState = *resp.Cluster.Status.State

		fmt.Println("Cluster State Is: " + clusterState)

		if clusterState == "WAITING" {
			_, ter_err := svc.TerminateJobFlows(&emr.TerminateJobFlowsInput{JobFlowIDs: []*string{clusterId}})
			if ter_err != nil {
				panic(ter_err)
			}

			clusterState = *resp.Cluster.Status.State
		}
		if clusterState == "TERMINATED" {
			break
		}
		fmt.Println("Waiting for 60 seconds before checking again")
		time.Sleep(60 * time.Second)
	}
}