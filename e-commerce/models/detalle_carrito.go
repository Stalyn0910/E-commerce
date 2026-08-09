package models

//==========================================================
// CLASE: DetalleCarrito
//==========================================================

type DetalleCarrito struct {

	id int

	producto *Producto

	cantidad int

	subtotal float64

}

// Constructor

func NuevoDetalleCarrito(

	id int,

	producto *Producto,

	cantidad int,

) *DetalleCarrito {

	d := &DetalleCarrito{

		id: id,

		producto: producto,

		cantidad: cantidad,

	}

	d.CalcularSubtotal()

	return d

}
func (d *DetalleCarrito) GetID() int {

	return d.id

}

func (d *DetalleCarrito) CalcularSubtotal() {

	d.subtotal = float64(d.cantidad) * d.producto.GetPrecio()

}

func (d *DetalleCarrito) GetSubtotal() float64 {

	return d.subtotal

}

func (d *DetalleCarrito) GetProducto() *Producto {

	return d.producto

}

func (d *DetalleCarrito) GetCantidad() int {

	return d.cantidad

}