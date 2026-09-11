package services

import "clientes/repositories"

type ClienteService struct {
	repository *repositories.ClienteRepository
}

func NewClienteService(repository *repositories.ClienteRepository) *ClienteService {
	return &ClienteService{repository: repository}
}

func (s *ClienteService) Crear(nombre, email string) repositories.Cliente {
	return s.repository.Crear(nombre, email)
}

func (s *ClienteService) ObtenerPorID(id string) (repositories.Cliente, error) {
	return s.repository.ObtenerPorID(id)
}
