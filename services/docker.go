package services

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// GetDockerInfo retourne les conteneurs et les images Docker
func GetDockerInfo() (map[string]interface{}, error) {
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

	info := map[string]interface{}{
		"containers": containers,
		"images":     images,
	}

	return info, nil
}
