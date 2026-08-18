package services

import (
	"errors"
	"e-commerce/models"
	"sync"
)

//==========================================================
// SERVICIO: ProductoService
//==========================================================

type ProductoService struct {

	productos map[int]*models.Producto
    siguienteID int
	mutex sync.RWMutex
}

//==========================================================
// Constructor
//==========================================================

func (s *ProductoService) GenerarID() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.siguienteID

	s.siguienteID++

	return id
}

func NuevoProductoService() *ProductoService {

	return &ProductoService{

		productos: make(map[int]*models.Producto),
        siguienteID: 1,
	}
}

//==========================================================
// Crear
//==========================================================
func (s *ProductoService) Crear(

	producto *models.Producto,

) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, existe := s.productos[producto.GetID()]; existe {

		return errors.New("producto existente")

	}

	s.productos[producto.GetID()] = producto

	return nil

}

//==========================================================
// Buscar
//==========================================================
func (s *ProductoService) Buscar(

	id int,

) (*models.Producto, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	producto, existe := s.productos[id]

	if !existe {

		return nil, errors.New("producto inexistente")

	}

	return producto, nil

}

//==========================================================
// Actualizar
//==========================================================
func (s *ProductoService) Actualizar(

	producto *models.Producto,

) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.productos[producto.GetID()] = producto

	return nil

}

//==========================================================
// Eliminar
//==========================================================
func (s *ProductoService) Eliminar(

	id int,

) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.productos, id)

	return nil

}

//==========================================================
// Listar
//==========================================================
func (s *ProductoService) Listar() []*models.Producto {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	lista := []*models.Producto{}

	for _, producto := range s.productos {

		lista = append(lista, producto)

	}

	return lista

}