package models

//==========================================================
// CLASE: CarritoCompras
//==========================================================

type CarritoCompras struct {

	id int

	detalles []*DetalleCarrito

	total float64

}

// Constructor

func NuevoCarrito(id int) *CarritoCompras {

	return &CarritoCompras{

		id: id,

		detalles: []*DetalleCarrito{},

	}

}

func (c *CarritoCompras) AgregarDetalle(

	detalle *DetalleCarrito,

) {

	c.detalles = append(

		c.detalles,

		detalle,

	)

}

func (c *CarritoCompras) SetDetalles(

	detalles []*DetalleCarrito,

) {

	c.detalles = detalles

}

func (c *CarritoCompras) GetDetalles() []*DetalleCarrito {

	return c.detalles

}

func (c *CarritoCompras) CalcularTotal() {

	total := 0.0

	for _, detalle := range c.detalles {

		total += detalle.GetSubtotal()

	}

	c.total = total

}

func (c *CarritoCompras) GetTotal() float64 {

	return c.total

}