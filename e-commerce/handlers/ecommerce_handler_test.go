package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"e-commerce/models"
	"e-commerce/services"
)

func TestCrearUsuarioAPI(t *testing.T) {

	service := services.NuevoUsuarioService()

	handler := NuevoUsuarioHandler(service)

	body := `{
		"nombre": "Stalyn",
		"email": "stalyn@gmail.com",
		"contrasena": "123456",
		"rol": "Cliente"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/usuarios",
		strings.NewReader(body),
	)

	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler.Crear(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"se esperaba código 201, se obtuvo %d",
			recorder.Code,
		)
	}
}

//==========================================================
// PRUEBA 5
// POST /api/carrito/productos
//
// Verifica que se pueda agregar un producto al carrito.
//==========================================================



func TestAgregarProductoCarritoAPI(t *testing.T) {

	productoService := services.NuevoProductoService()

	carritoService := services.NuevoCarritoService(1)

	producto := models.NuevoProducto(
		1,
		"Laptop",
		"Laptop Lenovo",
		850,
		10,
		"Tecnologia",
	)

	err := productoService.Crear(producto)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el producto: %v",
			err,
		)
	}

	handler := NuevoCarritoHandler(
		carritoService,
		productoService,
	)

	body := `{
		"producto_id": 1,
		"cantidad": 2
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/carrito/productos",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	handler.AgregarProducto(
		recorder,
		request,
	)

	if recorder.Code != http.StatusOK {

		t.Fatalf(
			"se esperaba código 200, se obtuvo %d",
			recorder.Code,
		)

	}
}

//==========================================================
// PRUEBA 6
// GET /api/carrito
//
// Verifica que se pueda consultar el carrito.
//==========================================================

func TestObtenerCarritoAPI(t *testing.T) {

	productoService := services.NuevoProductoService()

	carritoService := services.NuevoCarritoService(1)

	producto := models.NuevoProducto(
		1,
		"Mouse",
		"Mouse Logitech",
		25,
		20,
		"Accesorios",
	)

	err := productoService.Crear(producto)

	if err != nil {

		t.Fatalf(
			"no se pudo crear el producto: %v",
			err,
		)

	}

	err = carritoService.AgregarProducto(
		producto,
		2,
	)

	if err != nil {

		t.Fatalf(
			"no se pudo agregar el producto al carrito: %v",
			err,
		)

	}

	handler := NuevoCarritoHandler(
		carritoService,
		productoService,
	)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/carrito",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ObtenerCarrito(
		recorder,
		request,
	)

	if recorder.Code != http.StatusOK {

		t.Fatalf(
			"se esperaba código 200, se obtuvo %d",
			recorder.Code,
		)

	}
}

//==========================================================
// PRUEBA 7
// POST /api/pedidos
//
// Verifica que se pueda crear un pedido.
//==========================================================

func TestCrearPedidoAPI(t *testing.T) {

	usuarioService := services.NuevoUsuarioService()

	pedidoService := services.NuevoPedidoService()

	usuario := models.NuevoUsuario(
		1,
		"Stalyn",
		"stalyn@gmail.com",
		"123456",
		"Cliente",
	)

	err := usuarioService.Crear(usuario)

	if err != nil {

		t.Fatalf(
			"no se pudo crear el usuario: %v",
			err,
		)

	}

	handler := NuevoPedidoHandler(
		pedidoService,
		usuarioService,
	)

	body := `{
		"usuario_id": 1
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/pedidos",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	handler.Crear(
		recorder,
		request,
	)

	if recorder.Code != http.StatusCreated {

		t.Fatalf(
			"se esperaba código 201, se obtuvo %d",
			recorder.Code,
		)

	}
}

//==========================================================
// PRUEBA 8
// POST /api/pagos
//
// Verifica que se pueda registrar un pago.
//==========================================================

func TestRegistrarPagoAPI(t *testing.T) {

	pagoService := services.NuevoPagoService()

	handler := NuevoPagoHandler(
		pagoService,
	)

	body := `{
		"metodo_pago": "Tarjeta",
		"monto": 850
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/pagos",
		strings.NewReader(body),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	handler.Registrar(
		recorder,
		request,
	)

	if recorder.Code != http.StatusCreated {

		t.Fatalf(
			"se esperaba código 201, se obtuvo %d",
			recorder.Code,
		)

	}
}