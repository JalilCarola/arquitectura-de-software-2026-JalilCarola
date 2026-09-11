package services

import (
	"log"
	"sync"
	"time"

	"pedidos/repositories"
)

const cacheTTL = 10 * time.Second

// ProductoService agrega una caché simple en memoria delante del repositorio:
// mientras la caché sea válida, evita recalcular/releer el listado de productos.
type ProductoService struct {
	repository *repositories.ProductoRepository

	mu          sync.Mutex
	cache       []repositories.Producto
	cacheAt     time.Time
	cacheValido bool
}

func NewProductoService(repository *repositories.ProductoRepository) *ProductoService {
	return &ProductoService{repository: repository}
}

func (s *ProductoService) ObtenerTodos() []repositories.Producto {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cacheValido && time.Since(s.cacheAt) < cacheTTL {
		log.Println("[CACHE] Listado de productos servido desde caché.")
		return s.cache
	}

	log.Println("[CACHE] Caché vencida o vacía. Consultando el repositorio.")
	productos := s.repository.ObtenerTodos()
	s.cache = productos
	s.cacheAt = time.Now()
	s.cacheValido = true

	return productos
}

func (s *ProductoService) Existe(id string) bool {
	return s.repository.Existe(id)
}
