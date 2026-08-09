package models

//==========================================================
//
// CLASE: DetallePedido
//
// Representa un producto incluido dentro de un pedido.
//
//==========================================================

type DetallePedido struct {

	id int

	producto *Producto

	cantidad int

	subtotal float64

}

//==========================================================
// Constructor
//==========================================================

func NuevoDetallePedido(

	id int,

	producto *Producto,

	cantidad int,

) *DetallePedido {

	d := &DetallePedido{

		id: id,

		producto: producto,

		cantidad: cantidad,

	}

	d.CalcularSubtotal()

	return d

}

//==========================================================
// Calcula el subtotal.
//==========================================================

func (d *DetallePedido) CalcularSubtotal() {

	d.subtotal = float64(d.cantidad) * d.producto.GetPrecio()

}

//==========================================================
// GETTERS
//==========================================================

func (d *DetallePedido) GetID() int {

	return d.id

}

func (d *DetallePedido) GetProducto() *Producto {

	return d.producto

}

func (d *DetallePedido) GetCantidad() int {

	return d.cantidad

}

func (d *DetallePedido) GetSubtotal() float64 {

	return d.subtotal

}