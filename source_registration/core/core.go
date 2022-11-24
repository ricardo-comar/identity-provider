package core

import (
	"github.com/ricardo-comar/identity-provider/config"
)

type Core struct {
	Config *config.Config
}

func New() (*Core, error) {
	core := &Core{}
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	core.Config = cfg
	if err != nil {
		return nil, err
	}

	return core, nil
}
