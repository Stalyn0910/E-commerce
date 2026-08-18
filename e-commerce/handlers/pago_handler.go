package handlers

import (
	"encoding/json"
	"net/http"

	"e-commerce/models"
	"e-commerce/services"
)

//==========================================================
// HANDLER: Pagos
//==========================================================

type PagoHandler struct {
	service *services.PagoService
}

//==========================================================
// Constructor
//==========================================================

func NuevoPagoHandler(
	service *services.PagoService,
) *PagoHandler {

	return &PagoHandler{
		service: service,
	}
}

//==========================================================
// POST /api/pagos
//
// Registra un pago.
//==========================================================

func (h *PagoHandler) Registrar(
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
		MetodoPago string  `json:"metodo_pago"`
		Monto      float64 `json:"monto"`
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

	pago := models.NuevoPago(
		id,
		datos.MetodoPago,
		datos.Monto,
	)

	h.service.Registrar(pago)

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	w.WriteHeader(http.StatusCreated)

	respuesta := map[string]interface{}{
		"mensaje": "Pago registrado correctamente",
		"id":      pago.GetID(),
		"estado":  "Pagado",
	}

	json.NewEncoder(w).Encode(respuesta)
}
