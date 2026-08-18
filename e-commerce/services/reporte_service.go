package services

import (
	"fmt"

	"e-commerce/models"
)

//==========================================================
// SERVICIO: ReporteService
//
// Genera los reportes del sistema.
//
//==========================================================

type ReporteService struct{}

//==========================================================
// Reporte de usuarios
//==========================================================

func (r ReporteService) GenerarReporteUsuarios(

	usuarios []*models.Usuario,

) {

	fmt.Println("========== USUARIOS ==========")

	for _, usuario := range usuarios {

		fmt.Printf(

			"%d - %s - %s\n",

			usuario.GetID(),

			usuario.GetNombre(),

			usuario.GetEmail(),

		)

	}

}

//==========================================================
// Reporte de productos
//==========================================================

func (r ReporteService) GenerarReporteProductos(

	productos []*models.Producto,

) {

	fmt.Println("========== PRODUCTOS ==========")

	for _, producto := range productos {

		fmt.Printf(

			"%d - %s - %.2f\n",

			producto.GetID(),

			producto.GetNombre(),

			producto.GetPrecio(),

		)

	}

}

//==========================================================
// Reporte de ventas
//==========================================================

func (r ReporteService) GenerarReporteVentas(

	pedidos []*models.Pedido,

) {

	fmt.Println()

	fmt.Println("========== REPORTE DE VENTAS ==========")

	totalGeneral := 0.0

	for _, pedido := range pedidos {

		fmt.Println("--------------------------------")

		fmt.Println("Pedido:", pedido.GetID())

		fmt.Println("Cliente:", pedido.GetUsuario().GetNombre())

		fmt.Printf("Total: %.2f\n", pedido.GetTotal())

		totalGeneral += pedido.GetTotal()

	}

	fmt.Println("--------------------------------")

	fmt.Printf("VENTAS TOTALES: %.2f\n", totalGeneral)

}