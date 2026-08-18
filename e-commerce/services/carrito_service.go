package services

import (
	"errors"
	"e-commerce/models"
)

//==========================================================
// SERVICIO: CarritoService
//==========================================================

type CarritoService struct {

	carrito *models.CarritoCompras

}

func NuevoCarritoService(

	id int,

) *CarritoService {

	return &CarritoService{

		carrito: models.NuevoCarrito(id),

	}

}

func (c *CarritoService) AgregarProducto(

	producto *models.Producto,

	cantidad int,

) error {

	if cantidad <= 0 {

		return errors.New("cantidad inválida")

	}

	if producto.GetStock() < cantidad {

		return errors.New("stock insuficiente")

	}

	detalle := models.NuevoDetalleCarrito(

		len(c.carrito.GetDetalles())+1,

		producto,

		cantidad,

	)

	c.carrito.AgregarDetalle(detalle)

	c.carrito.CalcularTotal()

	return nil

}

func (c *CarritoService) ObtenerCarrito() *models.CarritoCompras {

	return c.carrito

}