package services

import (
	"github.com/moby/moby/client"
)

func NewDockerClient(host string) (*client.Client, error) {
	return client.NewClientWithOpts(
		client.WithHost(host),
		client.FromEnv,
	)
}
