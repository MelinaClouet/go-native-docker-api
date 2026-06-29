package services

import (
	"context"
	"fmt"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func DeployProject(project models.Project) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	for _, service := range project.Services {

		resp, err := cli.ContainerCreate(
			context.Background(),
			&container.Config{
				Image: service.Image,
				Env:   mapToSlice(service.Environment),
			},
			nil, nil, nil,
			service.Name,
		)

		if err != nil {
			return err
		}

		if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
			return err
		}

		fmt.Println("Started:", service.Name)
	}

	return nil
}

func mapToSlice(env map[string]string) []string {
	var result []string
	for k, v := range env {
		result = append(result, k+"="+v)
	}
	return result
}
