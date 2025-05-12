package di

import (
	"errors"
	"log"

	"go.uber.org/dig"
)

type Repository struct{}
type Service struct {
	repo *Repository
}

func RunDI() {
	container := dig.New()
	container.Provide(newRepository)
	container.Provide(newService)
	if err := container.Invoke(checkService); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func newRepository() *Repository {
	return &Repository{}
}

func newService(repo *Repository) *Service {
	return &Service{
		repo,
	}
}

func checkService(s *Service) error {
	if s == nil {
		return errors.New("Service was not injected")
	}
	if s.repo != nil {
		log.Println("Repository was injected successfully")
	}
	log.Println("Service was injected successfully")
	return nil
}
