package services

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type DockerInfo struct {
	Containers []container.Summary `json:"containers"`
	Images     []image.Summary     `json:"images"`
}

func GetDockerInfo() (*DockerInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		return nil, err
	}

	return &DockerInfo{
		Containers: containers,
		Images:     images,
	}, nil
}
