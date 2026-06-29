package services

import (
	"context"

	"github.com/moby/moby/client"
)

type DockerInfo struct {
	Containers client.ContainerListResult
	Images     client.ImageListResult
}

func GetDockerInfo() (*DockerInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(context.Background(), client.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	return &DockerInfo{
		Containers: containers,
		Images:     images,
	}, nil
}

func GetDeploymentStatus(name string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	result, err := cli.ContainerList(context.Background(), client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return "", err
	}

	// Parcours des containers
	for _, ctr := range result.Items {
		for _, n := range ctr.Names {
			if n == "/"+name {
				return ctr.Status, nil
			}
		}
	}

	return "not-found", nil
}
