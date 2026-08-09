package interfaces

//==============================================================
// INTERFAZ: Reporte
//
// Define los reportes que puede generar el sistema.
//
//==============================================================

type Reporte interface {

	GenerarReporteUsuarios()

	GenerarReporteProductos()

	GenerarReporteVentas()

}