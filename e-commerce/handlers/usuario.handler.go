package handlers

import (
	"encoding/json"
	"net/http"

	"e-commerce/models"
	"e-commerce/services"
)

//==========================================================
// HANDLER: Usuarios
//==========================================================

type UsuarioHandler struct {
	service *services.UsuarioService
}

//==========================================================
// Constructor
//==========================================================

func NuevoUsuarioHandler(
	service *services.UsuarioService,
) *UsuarioHandler {

	return &UsuarioHandler{
		service: service,
	}
}

//==========================================================
// POST /api/usuarios
//
// Registra un nuevo usuario.
//==========================================================

func (h *UsuarioHandler) Crear(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var datos struct {
		Nombre     string `json:"nombre"`
		Email      string `json:"email"`
		Contrasena string `json:"contrasena"`
		Rol        string `json:"rol"`
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
	
	usuario := models.NuevoUsuario(
		id,
		datos.Nombre,
		datos.Email,
		datos.Contrasena,
		datos.Rol,
	)

	err = h.service.Crear(usuario)

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
		"mensaje": "Usuario registrado correctamente",
		"id":      usuario.GetID(),
	}

	json.NewEncoder(w).Encode(respuesta)
}

//==========================================================
// GET /api/usuarios
//
// Lista todos los usuarios.
//==========================================================

func (h *UsuarioHandler) Listar(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	usuarios := h.service.Listar()

	w.Header().Set(
		"Content-Type",
		"application/json; charset=utf-8",
	)

	json.NewEncoder(w).Encode(usuarios)
}