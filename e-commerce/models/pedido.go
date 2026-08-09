package models

import "time"

//==========================================================
//
// CLASE: Pedido
//
//==========================================================

type Pedido struct {

	id int

	fechaPedido time.Time

	estado string

	total float64

	usuario *Usuario

	pago *Pago

	detalles []*DetallePedido

}

//==========================================================
// Constructor
//==========================================================

func NuevoPedido(

	id int,

	usuario *Usuario,

) *Pedido {

	return &Pedido{

		id: id,

		usuario: usuario,

		fechaPedido: time.Now(),

		estado: "Pendiente",

		detalles: []*DetallePedido{},

	}

}

//==========================================================
// Agrega un detalle al pedido.
//==========================================================

func (p *Pedido) AgregarDetalle(

	detalle *DetallePedido,

) {

	p.detalles = append(

		p.detalles,

		detalle,

	)

	p.CalcularTotal()

}

//==========================================================
// Calcula el total.
//==========================================================

func (p *Pedido) CalcularTotal() {

	total := 0.0

	for _, detalle := range p.detalles {

		total += detalle.GetSubtotal()

	}

	p.total = total

}

//==========================================================
// GETTERS
//==========================================================

func (p *Pedido) GetID() int {

	return p.id

}

func (p *Pedido) GetEstado() string {

	return p.estado

}

func (p *Pedido) GetTotal() float64 {

	return p.total

}

func (p *Pedido) GetDetalles() []*DetallePedido {

	return p.detalles

}

func (p *Pedido) GetUsuario() *Usuario {

	return p.usuario

}

//==========================================================
// SETTERS
//==========================================================

func (p *Pedido) ActualizarEstado(

	estado string,

) {

	p.estado = estado

}