package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"p/models"
	"p/services"
)

//==========================================================
//
// SISTEMA DE GESTIÓN E-COMMERCE
//
// Proyecto desarrollado en Go.
//
// Basado en:
// • Programación Orientada a Objetos
// • Encapsulación
// • Interfaces
// • Constructores
// • Servicios
// • Manejo de errores
//
// Autor:
// Stalyn Guacales
//
//==========================================================

//==========================================================
// VARIABLES GLOBALES
//==========================================================

// Lector para capturar datos desde teclado.
var lector = bufio.NewReader(os.Stdin)

// Servicios del sistema.
var usuarioService = services.NuevoUsuarioService()

var productoService = services.NuevoProductoService()

var carritoService = services.NuevoCarritoService(1)

var pedidoService = services.NuevoPedidoService()

var pagoService = services.NuevoPagoService()

var reporteService = services.ReporteService{}

//==============================================
// Contadores automáticos para IDs
//==============================================

var siguienteUsuarioID = 1

var siguienteProductoID = 1

var siguientePedidoID = 1

var siguientePagoID = 1

//==========================================================
// FUNCIÓN PRINCIPAL
//==========================================================

func main() {

	for {

		fmt.Println()
		fmt.Println("==============================================")
		fmt.Println("      SISTEMA DE GESTIÓN E-COMMERCE")
		fmt.Println("==============================================")
		fmt.Println()

		fmt.Println("============== USUARIOS ==============")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Listar usuarios")
		fmt.Println("3. Buscar usuario")
		fmt.Println("4. Eliminar usuario")
		fmt.Println()

		fmt.Println("============= PRODUCTOS =============")
		fmt.Println("5. Registrar producto")
		fmt.Println("6. Listar productos")
		fmt.Println("7. Buscar producto")
		fmt.Println("8. Eliminar producto")
		fmt.Println()

		fmt.Println("============== CARRITO ==============")
		fmt.Println("9. Agregar producto al carrito")
		fmt.Println("10. Mostrar carrito")
		fmt.Println()

		fmt.Println("============== PEDIDOS ==============")
		fmt.Println("11. Crear pedido")
		fmt.Println()

		fmt.Println("=============== PAGOS ===============")
		fmt.Println("12. Registrar pago")
		fmt.Println()

		fmt.Println("============= REPORTES ==============")
		fmt.Println("13. Reporte de usuarios")
		fmt.Println("14. Reporte de productos")
		fmt.Println("15. Reporte de ventas")
		fmt.Println()

		fmt.Println("0. Salir")
		fmt.Println()

		opcion := leerEntero("Seleccione una opción: ")

		switch opcion {

		case 1:
			registrarUsuario()

		case 2:
			listarUsuarios()

		case 3:
			buscarUsuario()

		case 4:
			eliminarUsuario()

		case 5:
			registrarProducto()

		case 6:
			listarProductos()

		case 7:
			buscarProducto()

		case 8:
			eliminarProducto()

		case 9:
			agregarProductoCarrito()

		case 10:
			mostrarCarrito()

		case 11:
			crearPedido()

		case 12:
			registrarPago()

		case 13:
			reporteService.GenerarReporteUsuarios(
				usuarioService.Listar(),
			)

		case 14:
			reporteService.GenerarReporteProductos(
				productoService.Listar(),
			)

		case 15:
			reporteService.GenerarReporteVentas(
				pedidoService.Listar(),
			)

		case 0:

			fmt.Println()
			fmt.Println("Gracias por utilizar el sistema.")
			return

		default:

			fmt.Println()
			fmt.Println("Opción inválida.")

		}

		pausa()

	}

}

//==========================================================
// REGISTRAR USUARIO
//==========================================================

func registrarUsuario() {

	fmt.Println()
	fmt.Println("========== REGISTRAR USUARIO ==========")

	id := siguienteUsuarioID

	siguienteUsuarioID++

	nombre := leerCadena("Ingrese el nombre: ")

	email := leerCadena("Ingrese el correo: ")

	contrasena := leerCadena("Ingrese la contraseña: ")

	rol := leerCadena("Ingrese el rol (Administrador/Cliente): ")

	usuario := models.NuevoUsuario(

		id,

		nombre,

		email,

		contrasena,

		rol,
	)

	err := usuarioService.Crear(usuario)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Usuario registrado correctamente.")

	fmt.Println("ID generado:", usuario.GetID())
}

//==========================================================
// LISTAR USUARIOS
//==========================================================

func listarUsuarios() {

	fmt.Println()

	fmt.Println("========== LISTA DE USUARIOS ==========")

	usuarios := usuarioService.Listar()

	if len(usuarios) == 0 {

		fmt.Println("No existen usuarios registrados.")

		return

	}

	for _, usuario := range usuarios {

		fmt.Println("--------------------------------------")

		fmt.Println("ID:", usuario.GetID())

		fmt.Println("Nombre:", usuario.GetNombre())

		fmt.Println("Correo:", usuario.GetEmail())

		fmt.Println("Rol:", usuario.GetRol())

	}

}

//==========================================================
// BUSCAR USUARIO
//==========================================================

func buscarUsuario() {

	fmt.Println()

	fmt.Println("========== BUSCAR USUARIO ==========")

	id := leerEntero("Ingrese el ID del usuario: ")

	usuario, err := usuarioService.Buscar(id)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Usuario encontrado")

	fmt.Println("--------------------------------------")

	fmt.Println("ID:", usuario.GetID())

	fmt.Println("Nombre:", usuario.GetNombre())

	fmt.Println("Correo:", usuario.GetEmail())

	fmt.Println("Rol:", usuario.GetRol())

}

//==========================================================
// ELIMINAR USUARIO
//==========================================================

func eliminarUsuario() {

	fmt.Println()

	fmt.Println("========== ELIMINAR USUARIO ==========")

	id := leerEntero("Ingrese el ID del usuario: ")

	err := usuarioService.Eliminar(id)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Usuario eliminado correctamente.")

}

//==========================================================
// ACTUALIZAR USUARIO (Opcional)
//==========================================================

func actualizarUsuario() {

	fmt.Println()

	fmt.Println("========== ACTUALIZAR USUARIO ==========")

	id := leerEntero("Ingrese el ID del usuario: ")

	usuario, err := usuarioService.Buscar(id)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Deje el campo vacío si no desea modificarlo.")

	nombre := leerCadena("Nuevo nombre: ")

	if nombre != "" {

		if err := usuario.SetNombre(nombre); err != nil {

			fmt.Println(err)

			return

		}

	}

	email := leerCadena("Nuevo correo: ")

	if email != "" {

		if err := usuario.SetEmail(email); err != nil {

			fmt.Println(err)

			return

		}

	}

	contrasena := leerCadena("Nueva contraseña: ")

	if contrasena != "" {

		if err := usuario.SetContrasena(contrasena); err != nil {

			fmt.Println(err)

			return

		}

	}

	err = usuarioService.Actualizar(usuario)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Usuario actualizado correctamente.")

}

//==========================================================
// REGISTRAR PRODUCTO
//==========================================================

func registrarProducto() {

	fmt.Println()
	fmt.Println("========== REGISTRAR PRODUCTO ==========")

	id := siguienteProductoID

	siguienteProductoID++

	nombre := leerCadena("Ingrese el nombre: ")

	descripcion := leerCadena("Ingrese la descripción: ")

	precio := leerDecimal("Ingrese el precio: ")

	stock := leerEntero("Ingrese el stock: ")

	categoria := leerCadena("Ingrese la categoría: ")

	producto := models.NuevoProducto(

		id,

		nombre,

		descripcion,

		precio,

		stock,

		categoria,
	)

	err := productoService.Crear(producto)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Producto registrado correctamente.")

	fmt.Println("ID generado:", producto.GetID())

}

//==========================================================
// LISTAR PRODUCTOS
//==========================================================

func listarProductos() {

	fmt.Println()

	fmt.Println("========== LISTA DE PRODUCTOS ==========")

	productos := productoService.Listar()

	if len(productos) == 0 {

		fmt.Println("No existen productos registrados.")

		return

	}

	for _, producto := range productos {

		fmt.Println("----------------------------------------")

		fmt.Println("ID:", producto.GetID())

		fmt.Println("Nombre:", producto.GetNombre())

		fmt.Println("Descripción:", producto.GetDescripcion())

		fmt.Printf("Precio: %.2f\n", producto.GetPrecio())

		fmt.Println("Stock:", producto.GetStock())

		fmt.Println("Categoría:", producto.GetCategoria())

	}

}

//==========================================================
// BUSCAR PRODUCTO
//==========================================================

func buscarProducto() {

	fmt.Println()

	fmt.Println("========== BUSCAR PRODUCTO ==========")

	id := leerEntero("Ingrese el ID del producto: ")

	producto, err := productoService.Buscar(id)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Producto encontrado")

	fmt.Println("----------------------------------------")

	fmt.Println("ID:", producto.GetID())

	fmt.Println("Nombre:", producto.GetNombre())

	fmt.Println("Descripción:", producto.GetDescripcion())

	fmt.Printf("Precio: %.2f\n", producto.GetPrecio())

	fmt.Println("Stock:", producto.GetStock())

	fmt.Println("Categoría:", producto.GetCategoria())

}

//==========================================================
// ELIMINAR PRODUCTO
//==========================================================

func eliminarProducto() {

	fmt.Println()

	fmt.Println("========== ELIMINAR PRODUCTO ==========")

	id := leerEntero("Ingrese el ID del producto: ")

	err := productoService.Eliminar(id)

	if err != nil {

		fmt.Println()

		fmt.Println("Error:", err)

		return

	}

	fmt.Println()

	fmt.Println("Producto eliminado correctamente.")

}

//==========================================================
// ACTUALIZAR PRODUCTO
//==========================================================

func actualizarProducto() {

	fmt.Println()

	fmt.Println("========== ACTUALIZAR PRODUCTO ==========")

	id := leerEntero("Ingrese el ID del producto: ")

	producto, err := productoService.Buscar(id)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Deje el valor en 0 si no desea modificarlo.")

	precio := leerDecimal("Nuevo precio: ")

	if precio > 0 {

		if err := producto.SetPrecio(precio); err != nil {

			fmt.Println(err)

			return

		}

	}

	stock := leerEntero("Nuevo stock: ")

	if stock >= 0 {

		if err := producto.SetStock(stock); err != nil {

			fmt.Println(err)

			return

		}

	}

	err = productoService.Actualizar(producto)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Producto actualizado correctamente.")

}

//==========================================================
// AGREGAR PRODUCTO AL CARRITO
//==========================================================

func agregarProductoCarrito() {

	fmt.Println()
	fmt.Println("========== AGREGAR PRODUCTO AL CARRITO ==========")

	// Verificar que existan productos registrados.
	if len(productoService.Listar()) == 0 {

		fmt.Println("No existen productos registrados.")

		return

	}

	// Mostrar productos disponibles.
	fmt.Println()
	fmt.Println("Productos disponibles:")

	listarProductos()

	// Solicitar información.
	idProducto := leerEntero("Ingrese el ID del producto: ")

	producto, err := productoService.Buscar(idProducto)

	if err != nil {

		fmt.Println(err)

		return

	}

	cantidad := leerEntero("Cantidad: ")

	err = carritoService.AgregarProducto(

		producto,

		cantidad,
	)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Producto agregado correctamente al carrito.")

	fmt.Printf("Total actual: %.2f\n",

		carritoService.ObtenerCarrito().GetTotal())

}

//==========================================================
// MOSTRAR CARRITO
//==========================================================

func mostrarCarrito() {

	fmt.Println()
	fmt.Println("========== CARRITO DE COMPRAS ==========")

	carrito := carritoService.ObtenerCarrito()

	detalles := carrito.GetDetalles()

	if len(detalles) == 0 {

		fmt.Println("El carrito está vacío.")

		return

	}

	for _, detalle := range detalles {

		fmt.Println("----------------------------------------")

		fmt.Println("Producto :", detalle.GetProducto().GetNombre())

		fmt.Println("Cantidad :", detalle.GetCantidad())

		fmt.Printf("Subtotal : %.2f\n",

			detalle.GetSubtotal())

	}

	fmt.Println("----------------------------------------")

	fmt.Printf("TOTAL DEL CARRITO: %.2f\n",

		carrito.GetTotal())

}

//==========================================================
// CREAR PEDIDO
//==========================================================

func crearPedido() {

	fmt.Println()
	fmt.Println("========== CREAR PEDIDO ==========")

	// Validar que existan usuarios.
	if len(usuarioService.Listar()) == 0 {

		fmt.Println("No existen usuarios registrados.")

		return

	}

	// Validar carrito.
	carrito := carritoService.ObtenerCarrito()

	if len(carrito.GetDetalles()) == 0 {

		fmt.Println("El carrito está vacío.")

		return

	}

	idPedido := siguientePedidoID

	siguientePedidoID++

	idUsuario := leerEntero("ID del usuario: ")

	usuario, err := usuarioService.Buscar(idUsuario)

	if err != nil {

		fmt.Println(err)

		return

	}

	pedido := models.NuevoPedido(

		idPedido,

		usuario,
	)

	//----------------------------------------------------
	// Copiar los productos del carrito al pedido
	//----------------------------------------------------

	for _, detalle := range carrito.GetDetalles() {

		nuevoDetalle := models.NuevoDetallePedido(

			detalle.GetID(),

			detalle.GetProducto(),

			detalle.GetCantidad(),
		)

		pedido.AgregarDetalle(

			nuevoDetalle,
		)

	}

	pedidoService.Crear(

		pedido,
	)

	fmt.Println()

	fmt.Println("Pedido registrado correctamente.")

	fmt.Println("ID del pedido:", pedido.GetID())

	fmt.Printf("Total del pedido: %.2f\n",

		pedido.GetTotal())

}

//==========================================================
// REGISTRAR PAGO
//==========================================================

func registrarPago() {

	fmt.Println()
	fmt.Println("========== REGISTRAR PAGO ==========")

	if len(pedidoService.Listar()) == 0 {

		fmt.Println("No existen pedidos registrados.")

		return

	}

	idPago := siguientePagoID

	siguientePagoID++

	metodo := leerCadena("Método de pago: ")

	monto := carritoService.ObtenerCarrito().GetTotal()

	pago := models.NuevoPago(

		idPago,

		metodo,

		monto,
	)

	pagoService.Registrar(

		pago,
	)

	fmt.Println()

	fmt.Println("Pago registrado correctamente.")

	fmt.Println("ID del pago:", pago.GetID())

	if pago.ValidarPago() {

		fmt.Println("Estado: PAGADO")

	}

}

//==========================================================
// LEER CADENA
//==========================================================

func leerCadena(mensaje string) string {

	fmt.Print(mensaje)

	texto, _ := lector.ReadString('\n')

	texto = strings.TrimSpace(texto)

	return texto

}

//==========================================================
// LEER ENTERO
//==========================================================

func leerEntero(mensaje string) int {

	for {

		fmt.Print(mensaje)

		texto, _ := lector.ReadString('\n')

		texto = strings.TrimSpace(texto)

		numero, err := strconv.Atoi(texto)

		if err == nil {

			return numero

		}

		fmt.Println("Ingrese un número válido.")

	}

}

//==========================================================
// LEER DECIMAL
//==========================================================

func leerDecimal(mensaje string) float64 {

	for {

		fmt.Print(mensaje)

		texto, _ := lector.ReadString('\n')

		texto = strings.TrimSpace(texto)

		numero, err := strconv.ParseFloat(texto, 64)

		if err == nil {

			return numero

		}

		fmt.Println("Ingrese un valor válido.")

	}

}

//==========================================================
// PAUSA
//==========================================================

func pausa() {

	fmt.Println()

	fmt.Print("Presione ENTER para continuar...")

	lector.ReadString('\n')

	fmt.Println()

}
