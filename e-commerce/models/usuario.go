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
	ID          int
	Nombre      string
	Email       string
	Contrasena  string
	Rol         string
}

//==========================================================
// Constructor
//==========================================================

func NuevoUsuario(id int, nombre, email, contrasena, rol string) *Usuario {
	return &Usuario{
		ID: id,
		Nombre: nombre,
		Email: email,
		Contrasena: contrasena,
		Rol: rol,
	}
}

//================ GETTERS =================

func (u *Usuario) GetID() int {
	return u.ID
}

func (u *Usuario) GetNombre() string {
	return u.Nombre
}

func (u *Usuario) GetEmail() string {
	return u.Email
}

func (u *Usuario) GetRol() string {
	return u.Rol
}

//================ SETTERS =================

func (u *Usuario) SetNombre(nombre string) error {

	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	u.Nombre = nombre
	return nil
}

func (u *Usuario) SetEmail(email string) error {

	if email == "" {
		return errors.New("el correo no puede estar vacío")
	}

	u.Email = email
	return nil
}

func (u *Usuario) SetContrasena(password string) error {

	if len(password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	u.Contrasena = password
	return nil
}