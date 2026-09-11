package repositories

import (
	"errors"
	"strconv"
	"sync"
)

var ErrClienteNoEncontrado = errors.New("cliente no encontrado")

type Cliente struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

type ClienteRepository struct {
	mu       sync.Mutex
	clientes map[string]Cliente
	nextID   int
}

func NewClienteRepository() *ClienteRepository {
	return &ClienteRepository{
		clientes: make(map[string]Cliente),
		nextID:   1,
	}
}

func (r *ClienteRepository) Crear(nombre, email string) Cliente {
	r.mu.Lock()
	defer r.mu.Unlock()

	cliente := Cliente{
		ID:     "C-" + strconv.Itoa(r.nextID),
		Nombre: nombre,
		Email:  email,
	}
	r.clientes[cliente.ID] = cliente
	r.nextID++

	return cliente
}

func (r *ClienteRepository) ObtenerPorID(id string) (Cliente, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cliente, ok := r.clientes[id]
	if !ok {
		return Cliente{}, ErrClienteNoEncontrado
	}

	return cliente, nil
}
