package models

import "errors"

//==========================================================
// CLASE: Producto
//==========================================================

type Producto struct {

	ID int

	Nombre string

	Descripcion string

	Precio float64

	Stock int

	Categoria string

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

		ID: id,

		Nombre: nombre,

		Descripcion: descripcion,

		Precio: precio,

		Stock: stock,

		Categoria: categoria,

	}

}

//================ GETTERS =================

func (p *Producto) GetID() int {
	return p.ID
}

func (p *Producto) GetNombre() string {
	return p.Nombre
}

func (p *Producto) GetDescripcion() string {
	return p.Descripcion
}

func (p *Producto) GetPrecio() float64 {
	return p.Precio
}

func (p *Producto) GetStock() int {
	return p.Stock
}

func (p *Producto) GetCategoria() string {
	return p.Categoria
}

//================ SETTERS =================

func (p *Producto) SetPrecio(precio float64) error {

	if precio <= 0 {
		return errors.New("precio inválido")
	}

	p.Precio = precio

	return nil
}

func (p *Producto) SetStock(stock int) error {

	if stock < 0 {
		return errors.New("stock inválido")
	}

	p.Stock = stock

	return nil
}