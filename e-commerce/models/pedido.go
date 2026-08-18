package models

import "time"

//==========================================================
//
// CLASE: Pedido
//
//==========================================================

type Pedido struct {

	ID int

	FechaPedido time.Time

	Estado string

	Total float64

	Usuario *Usuario

	Pago *Pago

	Detalles []*DetallePedido

}

//==========================================================
// Constructor
//==========================================================

func NuevoPedido(

	id int,

	usuario *Usuario,

) *Pedido {

	return &Pedido{

		ID: id,

		Usuario: usuario,

		FechaPedido: time.Now(),

		Estado: "Pendiente",

		Detalles: []*DetallePedido{},

	}

}

//==========================================================
// Agrega un detalle al pedido.
//==========================================================

func (p *Pedido) AgregarDetalle(

	detalle *DetallePedido,

) {

	p.Detalles = append(

		p.Detalles,

		detalle,

	)

	p.CalcularTotal()

}

//==========================================================
// Calcula el total.
//==========================================================

func (p *Pedido) CalcularTotal() {

	total := 0.0

	for _, detalle := range p.Detalles {

		total += detalle.GetSubtotal()

	}

	p.Total = total

}

//==========================================================
// GETTERS
//==========================================================

func (p *Pedido) GetID() int {

	return p.ID

}

func (p *Pedido) GetEstado() string {

	return p.Estado

}

func (p *Pedido) GetTotal() float64 {

	return p.Total

}

func (p *Pedido) GetDetalles() []*DetallePedido {

	return p.Detalles

}

func (p *Pedido) GetUsuario() *Usuario {

	return p.Usuario

}

//==========================================================
// SETTERS
//==========================================================

func (p *Pedido) ActualizarEstado(

	estado string,

) {

	p.Estado = estado

}