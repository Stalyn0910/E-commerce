package handlers

import (
	"encoding/json"
	"net/http"

	"e-commerce/models"
	"e-commerce/services"
)

//==========================================================
// HANDLER: Pedidos
//==========================================================

type PedidoHandler struct {
	service *services.PedidoService
	usuarioService *services.UsuarioService
}

//==========================================================
// Constructor
//==========================================================

func NuevoPedidoHandler(
	service *services.PedidoService,
	usuarioservice *services.UsuarioService,
) *PedidoHandler {

	return &PedidoHandler{
		service: service,
		usuarioService: usuarioservice,
	}
}

//==========================================================
// POST /api/pedidos
//
// Crea un nuevo pedido.
//==========================================================

func (h *PedidoHandler) Crear(
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
		UsuarioID int `json:"usuario_id"`
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

	usuario, err := h.usuarioService.Buscar(
	datos.UsuarioID,)

    if err != nil {
	    http.Error(
		    w,
		    err.Error(),
		    http.StatusNotFound,
	    )
	    return
    }

	id := h.service.GenerarID()

	pedido := models.NuevoPedido(
		id,
		usuario,
	)

	h.service.Crear(pedido)

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	w.WriteHeader(http.StatusCreated)

	respuesta := map[string]interface{}{
		"mensaje": "Pedido creado correctamente",
		"id":      pedido.GetID(),
		"estado":  pedido.GetEstado(),
	}

	json.NewEncoder(w).Encode(respuesta)
}