package services

import (
	"testing"

	"e-commerce/models"
)
func TestCarritoAgregarProducto(t *testing.T) {

	producto := models.NuevoProducto(
		1,
		"Laptop",
		"Laptop Lenovo",
		850,
		10,
		"Tecnología",
	)

	service := NuevoCarritoService(1)

	err := service.AgregarProducto(
		producto,
		2,
	)

	if err != nil {

		t.Fatalf(
			"no se pudo agregar el producto: %v",
			err,
		)
	}

	carrito := service.ObtenerCarrito()

	esperado := 1700.0

	if carrito.GetTotal() != esperado {

		t.Errorf(
			"total incorrecto: se esperaba %.2f, se obtuvo %.2f",
			esperado,
			carrito.GetTotal(),
		)
	}
}
func TestCarritoStockInsuficiente(t *testing.T) {

	producto := models.NuevoProducto(
		1,
		"Mouse",
		"Mouse Logitech",
		25,
		2,
		"Accesorios",
	)

	service := NuevoCarritoService(1)

	err := service.AgregarProducto(
		producto,
		5,
	)

	if err == nil {

		t.Error(
			"se permitió agregar una cantidad superior al stock",
		)
	}
}