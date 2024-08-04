package service

import (
	"app/internal/repository"
)

type Example struct {
	exampleRepo repository.IExampleRepo
}

func NewExampleService(exampleRepo repository.IExampleRepo) *Example {
	return &Example{exampleRepo: exampleRepo}
}

func (s *Example) Create(name string) {
	s.exampleRepo.Create(name)
}
