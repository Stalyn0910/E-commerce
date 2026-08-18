package services

import (
	"testing"

	"e-commerce/models"
)
func TestProductoServiceCrear(t *testing.T) {

	service := NuevoProductoService()

	producto := models.NuevoProducto(
		service.GenerarID(),
		"Laptop",
		"Laptop Lenovo",
		850,
		10,
		"Tecnología",
	)

	err := service.Crear(producto)

	if err != nil {

		t.Fatalf(
			"no se pudo crear el producto: %v",
			err,
		)
	}

	resultado, err := service.Buscar(producto.GetID())

	if err != nil {

		t.Fatalf(
			"no se encontró el producto: %v",
			err,
		)
	}

	if resultado.GetNombre() != "Laptop" {

		t.Errorf(
			"se esperaba Laptop, se obtuvo %s",
			resultado.GetNombre(),
		)
	}
}
