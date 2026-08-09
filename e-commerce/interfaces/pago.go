package interfaces

//==============================================================
// INTERFAZ: PagoInterface
//
// Define el comportamiento de cualquier método de pago.
//
//==============================================================

type PagoInterface interface {

	RegistrarPago()

	ValidarPago() bool

}