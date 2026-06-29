package services

import (
	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/client"
)

func NewDockerClient(host string) (*client.Client, error) {
	utils.Logger.Printf("[ENGINE] creating docker client host=%s", host)

	cli, err := client.NewClientWithOpts(
		client.WithHost(host),
		client.FromEnv,
	)

	if err != nil {
		utils.Logger.Printf("[ERROR][ENGINE] failed to create docker client host=%s err=%v", host, err)
		return nil, err
	}

	utils.Logger.Printf("[ENGINE] docker client created host=%s", host)

	return cli, nil
}
