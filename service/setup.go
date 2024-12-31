package service

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/hashicorp/nomad/api"
)

type Service struct {
	Validator   *validator.Validate
	NomadClient *api.Client
}

func Setup() *Service {
	client, err := api.NewClient(&api.Config{
		Address: "http://127.0.0.1:4646",
	})
	if err != nil {
		panic(fmt.Sprintf("Error creating Hashicorp Nomad client: %s", err))
	}

	return &Service{
		Validator:   validator.New(),
		NomadClient: client,
	}
}
