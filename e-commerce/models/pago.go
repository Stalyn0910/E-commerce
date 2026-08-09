package models

import "time"

type Pago struct {

	id int

	metodoPago string

	fechaPago time.Time

	monto float64

	estadoPago string

}

func NuevoPago(

	id int,

	metodo string,

	monto float64,

) *Pago {

	return &Pago{

		id: id,

		metodoPago: metodo,

		monto: monto,

		fechaPago: time.Now(),

		estadoPago: "Pendiente",

	}

}

func (p *Pago) GetID() int {

	return p.id

}

func (p *Pago) RegistrarPago() {

	p.estadoPago = "Pagado"

}

func (p *Pago) ValidarPago() bool {

	return p.estadoPago == "Pagado"

}