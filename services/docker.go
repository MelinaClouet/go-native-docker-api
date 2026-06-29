package services

import (
	"context"

	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/client"
)

type DockerInfo struct {
	Containers client.ContainerListResult
	Images     client.ImageListResult
}

func GetDockerInfo() (*DockerInfo, error) {
	utils.Logger.Println("[DOCKER_INFO] fetching docker info")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] docker client init failed: %v", err)
		return nil, err
	}

	utils.Logger.Println("[DOCKER_INFO] listing containers")

	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{})
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] container list failed: %v", err)
		return nil, err
	}

	utils.Logger.Printf("[DOCKER_INFO] containers found=%d", len(containers.Items))

	utils.Logger.Println("[DOCKER_INFO] listing images")

	images, err := cli.ImageList(context.Background(), client.ImageListOptions{})
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] image list failed: %v", err)
		return nil, err
	}

	utils.Logger.Printf("[DOCKER_INFO] images found=%d", len(images.Items))

	return &DockerInfo{
		Containers: containers,
		Images:     images,
	}, nil
}

func GetDeploymentStatus(name string) (string, error) {
	utils.Logger.Printf("[STATUS] checking container name=%s", name)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		utils.Logger.Printf("[ERROR][STATUS] docker client init failed: %v", err)
		return "", err
	}

	utils.Logger.Println("[STATUS] listing containers (All=true)")

	result, err := cli.ContainerList(context.Background(), client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		utils.Logger.Printf("[ERROR][STATUS] container list failed: %v", err)
		return "", err
	}

	utils.Logger.Printf("[STATUS] total containers=%d", len(result.Items))

	// Parcours des containers
	for _, ctr := range result.Items {
		for _, n := range ctr.Names {
			if n == "/"+name {
				utils.Logger.Printf("[STATUS] found container name=%s status=%s", name, ctr.Status)
				return ctr.Status, nil
			}
		}
	}

	utils.Logger.Printf("[STATUS] container not found name=%s", name)
	return "not-found", nil
}
