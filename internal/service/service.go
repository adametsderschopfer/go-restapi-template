package service

import "app/internal/repository"

type Deps struct {
	Repos *repository.Repositories
}

type Services struct {
	Example Example
}

func NewServices(deps Deps) *Services {
	exampleService := NewExampleService(deps.Repos.Example)

	return &Services{
		Example: *exampleService,
	}
}
