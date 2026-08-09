package services

import (
	"errors"
	"p/models"
)

//==========================================================
// SERVICIO: ProductoService
//==========================================================

type ProductoService struct {

	productos map[int]*models.Producto

}

func NuevoProductoService() *ProductoService {

	return &ProductoService{

		productos: make(map[int]*models.Producto),

	}

}

func (s *ProductoService) Crear(

	producto *models.Producto,

) error {

	if _, existe := s.productos[producto.GetID()]; existe {

		return errors.New("producto existente")

	}

	s.productos[producto.GetID()] = producto

	return nil

}

func (s *ProductoService) Buscar(

	id int,

) (*models.Producto, error) {

	producto, existe := s.productos[id]

	if !existe {

		return nil, errors.New("producto inexistente")

	}

	return producto, nil

}

func (s *ProductoService) Actualizar(

	producto *models.Producto,

) error {

	s.productos[producto.GetID()] = producto

	return nil

}

func (s *ProductoService) Eliminar(

	id int,

) error {

	delete(s.productos, id)

	return nil

}

func (s *ProductoService) Listar() []*models.Producto {

	lista := []*models.Producto{}

	for _, producto := range s.productos {

		lista = append(lista, producto)

	}

	return lista

}