package models

import "errors"

//==========================================================
// CLASE: Producto
//==========================================================

type Producto struct {

	id int

	nombre string

	descripcion string

	precio float64

	stock int

	categoria string

}

// Constructor

func NuevoProducto(

	id int,

	nombre string,

	descripcion string,

	precio float64,

	stock int,

	categoria string,

) *Producto {

	return &Producto{

		id: id,

		nombre: nombre,

		descripcion: descripcion,

		precio: precio,

		stock: stock,

		categoria: categoria,

	}

}

//================ GETTERS =================

func (p *Producto) GetID() int {
	return p.id
}

func (p *Producto) GetNombre() string {
	return p.nombre
}

func (p *Producto) GetDescripcion() string {
	return p.descripcion
}

func (p *Producto) GetPrecio() float64 {
	return p.precio
}

func (p *Producto) GetStock() int {
	return p.stock
}

func (p *Producto) GetCategoria() string {
	return p.categoria
}

//================ SETTERS =================

func (p *Producto) SetPrecio(precio float64) error {

	if precio <= 0 {
		return errors.New("precio inválido")
	}

	p.precio = precio

	return nil
}

func (p *Producto) SetStock(stock int) error {

	if stock < 0 {
		return errors.New("stock inválido")
	}

	p.stock = stock

	return nil
}