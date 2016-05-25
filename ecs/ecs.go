package ecs

import(
	"fmt"
	"os"

	// aws "github.com/aws/aws-sdk-go/aws"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_ecs "github.com/aws/aws-sdk-go/service/ecs"
)

// `ECSClient` is the polling instance around the ECS api that returns the services and tasks
// for a given cluster.
type ECSClient struct {
	client aws_ecs.ECS
}

//
type Task {
	// ip net.IP
	// ports []int
}

//
type Service {
	name string
	tasks []Task
}

//
func (c ECSClient) Refresh() ([]Service, error) {
	return nil, nil
}

// `NewClient` validates the ECS configuration and returns an ECS client instance
// for refreshing the services and tasks.
func NewClient() (*ECSClient, error) {
	// get the cluster name from AWS_ECS_CLUSTER_NAME
	// the other env vars required are passed through to the aws client
	clusterName := os.Getenv("AWS_ECS_CLUSTER_NAME")
	if clusterName == "" {
		return nil, fmt.Errorf("No cluster name defined, set `AWS_ECS_CLUSTER_NAME`")
	}

	ecsClient := aws_ecs.New(aws_session.New())
	if ecsClient == nil {
		return nil, fmt.Errorf("Error creating ECS client.")
	}

	return nil, nil
}

// params := &ecs.ListServicesInput{
// 	Cluster:    aws.String("String"),
// 	MaxResults: aws.Int64(1),
// 	NextToken:  aws.String("String"),
// }
// resp, err := svc.ListServices(params)
