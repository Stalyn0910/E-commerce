package handlers

import (
	"encoding/json"
	"net/http"
	"e-commerce/services"
)

//==========================================================
// HANDLER: Carrito
//==========================================================

type CarritoHandler struct {
	service *services.CarritoService

	productoService *services.ProductoService
}

//==========================================================
// Constructor
//==========================================================

func NuevoCarritoHandler(
	service *services.CarritoService,
	productoService *services.ProductoService,
) *CarritoHandler {

	return &CarritoHandler{
		service: service,

		productoService: productoService,
	}
}

//==========================================================
// POST /api/carrito/productos
//
// Agrega un producto al carrito.
//==========================================================

func (h *CarritoHandler) AgregarProducto(
	w http.ResponseWriter,
	r *http.Request,

) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var datos struct {
		ProductoID int `json:"producto_id"`
		Cantidad   int `json:"cantidad"`
	}

	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(
			w,
			"JSON inválido",
			http.StatusBadRequest,
		)
		return
	}

	producto, err := h.productoService.Buscar(
	datos.ProductoID,)

    if err != nil {
	    http.Error(
		    w,
		    err.Error(),
		    http.StatusNotFound,
	    )
	    return
    }

	err = h.service.AgregarProducto(
		producto,
		datos.Cantidad,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	respuesta := map[string]interface{}{
		"mensaje": "Producto agregado al carrito",
		"total":   h.service.ObtenerCarrito().GetTotal(),
	}

	json.NewEncoder(w).Encode(respuesta)
}

//==========================================================
// GET /api/carrito
//
// Consulta el carrito.
//==========================================================

func (h *CarritoHandler) ObtenerCarrito(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	carrito := h.service.ObtenerCarrito()

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	respuesta := map[string]interface{}{
		"total":    carrito.GetTotal(),
		"detalles": carrito.GetDetalles(),
	}

	json.NewEncoder(w).Encode(respuesta)
}