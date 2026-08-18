package handlers

import (
	"encoding/json"
	"net/http"

	"e-commerce/models"
	"e-commerce/services"
)

//==========================================================
// HANDLER: Productos
//==========================================================

type ProductoHandler struct {
	service *services.ProductoService
}

//==========================================================
// Constructor
//==========================================================

func NuevoProductoHandler(
	service *services.ProductoService,
) *ProductoHandler {

	return &ProductoHandler{
		service: service,
	}
}

//==========================================================
// POST /api/productos
//
// Registra un nuevo producto.
//==========================================================

func (h *ProductoHandler) Crear(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var datos struct {
		Nombre      string  `json:"nombre"`
		Descripcion string  `json:"descripcion"`
		Precio      float64 `json:"precio"`
		Stock       int     `json:"stock"`
		Categoria   string  `json:"categoria"`
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

	id := h.service.GenerarID()

	producto := models.NuevoProducto(
		id,
		datos.Nombre,
		datos.Descripcion,
		datos.Precio,
		datos.Stock,
		datos.Categoria,
	)

	err = h.service.Crear(producto)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusConflict,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	w.WriteHeader(http.StatusCreated)

	respuesta := map[string]interface{}{
		"mensaje": "Producto registrado correctamente",
		"id":      producto.GetID(),
	}

	json.NewEncoder(w).Encode(respuesta)
}

//==========================================================
// GET /api/productos
//
// Lista todos los productos.
//==========================================================

func (h *ProductoHandler) Listar(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	productos := h.service.Listar()

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	json.NewEncoder(w).Encode(productos)
}
