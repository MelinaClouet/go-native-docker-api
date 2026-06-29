package services

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func DeployProject(project models.Project) error {
	log.Println("Déploiement :", project.Name)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	for _, service := range project.Services {

		log.Println("Pull:", service.Image)

		reader, err := cli.ImagePull(ctx, service.Image, client.ImagePullOptions{})
		if err != nil {
			return err
		}

		_, err = io.Copy(io.Discard, reader)
		reader.Close()
		if err != nil {
			return err
		}

		resp, err := cli.ContainerCreate(
			ctx,
			client.ContainerCreateOptions{
				Config: &container.Config{
					Image: service.Image,
					Env:   utils.MapToSlice(service.Environment),
				},
				Name: service.Name,
			},
		)
		if err != nil {
			return err
		}

		log.Println("Container:", resp.ID)

		_, err = cli.ContainerStart(ctx, resp.ID, client.ContainerStartOptions{})
		if err != nil {
			return err
		}

		log.Println("Container démarré :", resp.ID)
	}

	return nil
}
