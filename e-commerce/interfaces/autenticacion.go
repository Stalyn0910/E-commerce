package interfaces

//==============================================================
// INTERFAZ: Autenticacion
//
// Define las operaciones que debe implementar cualquier
// entidad que pueda autenticarse.
//
//==============================================================

type Autenticacion interface {

	Registrar()

	IniciarSesion()

}