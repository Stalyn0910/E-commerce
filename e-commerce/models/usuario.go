package models

import (
	"errors"
)

//==========================================================
// CLASE: Usuario
//
// Representa un usuario del sistema.
//==========================================================

type Usuario struct {
	id          int
	nombre      string
	email       string
	contrasena  string
	rol         string
}

//==========================================================
// Constructor
//==========================================================

func NuevoUsuario(id int, nombre, email, contrasena, rol string) *Usuario {
	return &Usuario{
		id: id,
		nombre: nombre,
		email: email,
		contrasena: contrasena,
		rol: rol,
	}
}

//================ GETTERS =================

func (u *Usuario) GetID() int {
	return u.id
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) GetEmail() string {
	return u.email
}

func (u *Usuario) GetRol() string {
	return u.rol
}

//================ SETTERS =================

func (u *Usuario) SetNombre(nombre string) error {

	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	u.nombre = nombre
	return nil
}

func (u *Usuario) SetEmail(email string) error {

	if email == "" {
		return errors.New("el correo no puede estar vacío")
	}

	u.email = email
	return nil
}

func (u *Usuario) SetContrasena(password string) error {

	if len(password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	u.contrasena = password
	return nil
}