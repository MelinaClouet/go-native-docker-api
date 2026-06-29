package services

import (
	"context"
	"io"
	"time"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

// DeployProject déploie un projet complet sur un engine Docker cible.
//
// Le déploiement suit les étapes suivantes pour chaque service :
//  1. Connexion au Docker Engine (local ou distant)
//  2. Pull de l'image Docker
//  3. Création du conteneur
//  4. Démarrage du conteneur
//
// Le projet peut contenir plusieurs services exécutés indépendamment.
//
// Chaque étape est journalisée pour assurer la traçabilité et le debug.
func DeployProject(project models.Project) error {

	// Log de début de déploiement
	utils.Logger.Printf("[DEPLOY] start project=%s engine=%s",
		project.Name,
		project.Engine.Host,
	)

	// Connexion au Docker Engine cible (local ou distant)
	cli, err := NewDockerClient(project.Engine.Host)
	if err != nil {
		utils.Logger.Printf("[ERROR][ENGINE] project=%s err=%v",
			project.Name, err,
		)
		return err
	}

	// Timeout global du déploiement (10 minutes max)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Boucle sur tous les services du projet
	for _, service := range project.Services {

		utils.Logger.Printf("[SERVICE] project=%s service=%s image=%s",
			project.Name, service.Name, service.Image,
		)

		// -------------------------
		// PULL IMAGE DOCKER
		// -------------------------
		utils.Logger.Printf("[PULL] image=%s", service.Image)

		reader, err := cli.ImagePull(ctx, service.Image, client.ImagePullOptions{})
		if err != nil {
			utils.Logger.Printf("[ERROR][PULL] image=%s err=%v",
				service.Image, err,
			)
			return err
		}

		// Consommation du stream de pull
		_, err = io.Copy(io.Discard, reader)
		reader.Close()
		if err != nil {
			utils.Logger.Printf("[ERROR][PULL][READ] image=%s err=%v",
				service.Image, err,
			)
			return err
		}

		utils.Logger.Printf("[PULL] success image=%s", service.Image)

		// -------------------------
		// CREATE CONTAINER
		// -------------------------
		utils.Logger.Printf("[CONTAINER][CREATE] name=%s image=%s",
			service.Name, service.Image,
		)

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
			utils.Logger.Printf("[ERROR][CONTAINER][CREATE] service=%s err=%v",
				service.Name, err,
			)
			return err
		}

		utils.Logger.Printf("[CONTAINER][CREATED] name=%s id=%s",
			service.Name, resp.ID,
		)

		// -------------------------
		// START CONTAINER
		// -------------------------
		utils.Logger.Printf("[CONTAINER][START] id=%s", resp.ID)

		_, err = cli.ContainerStart(ctx, resp.ID, client.ContainerStartOptions{})
		if err != nil {
			utils.Logger.Printf("[ERROR][CONTAINER][START] id=%s err=%v",
				resp.ID, err,
			)
			return err
		}

		utils.Logger.Printf("[CONTAINER][RUNNING] id=%s", resp.ID)
	}

	// Fin du déploiement
	utils.Logger.Printf("[DEPLOY] success project=%s", project.Name)

	return nil
}
