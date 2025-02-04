package cmd

import (
	"RMazeE-server/config"
	"RMazeE-server/network"
	"RMazeE-server/repository"
	"RMazeE-server/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewCmd() *Cmd {

	c := &Cmd{
		config: config.NewConfig(),
	}
	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)
	c.network.ServerStart(c.config.Server.Port)
	return c
}
