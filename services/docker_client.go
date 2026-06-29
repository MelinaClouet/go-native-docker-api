package services

import (
	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/client"
)

// NewDockerClient crée une connexion vers un Docker Engine.
//
// Cette fonction permet de se connecter soit :
//   - à un daemon local (unix socket)
//   - à un daemon distant (TCP : remote VPS, serveur staging, etc.)
//
// Elle est utilisée dans la phase multi-engine pour cibler
// dynamiquement un serveur Docker en fonction du projet.
func NewDockerClient(host string) (*client.Client, error) {

	// Log de tentative de connexion
	utils.Logger.Printf("[ENGINE] creating docker client host=%s", host)

	// Création du client Docker avec configuration du host cible
	cli, err := client.NewClientWithOpts(
		client.WithHost(host),
		client.FromEnv,
	)

	// Gestion erreur de connexion
	if err != nil {
		utils.Logger.Printf("[ERROR][ENGINE] failed to create docker client host=%s err=%v",
			host, err,
		)
		return nil, err
	}

	// Log succès connexion engine
	utils.Logger.Printf("[ENGINE] docker client created host=%s", host)

	return cli, nil
}
