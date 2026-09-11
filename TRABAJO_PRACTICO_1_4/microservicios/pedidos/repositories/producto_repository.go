package repositories

type Producto struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Stock  int    `json:"stock"`
}

type ProductoRepository struct {
	productos []Producto
}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{
		productos: []Producto{
			{ID: "P-1", Nombre: "Auriculares", Stock: 10},
			{ID: "P-2", Nombre: "Teclado", Stock: 8},
		},
	}
}

func (r *ProductoRepository) ObtenerTodos() []Producto {
	result := make([]Producto, len(r.productos))
	copy(result, r.productos)
	return result
}

func (r *ProductoRepository) Existe(id string) bool {
	for _, producto := range r.productos {
		if producto.ID == id {
			return true
		}
	}
	return false
}
