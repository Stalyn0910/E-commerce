package interfaces

//==============================================================
// INTERFAZ GENÉRICA: Repository
//
// Permite reutilizar el mismo contrato para cualquier
// entidad del sistema.
//
//==============================================================

type Repository[T any] interface {

	Crear(objeto *T) error

	Actualizar(objeto *T) error

	Eliminar(id int) error

	Buscar(id int) (*T, error)

	Listar() []*T

}