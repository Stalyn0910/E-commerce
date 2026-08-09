package services

import (
	"errors"
	"p/models"
)

//==========================================================
// SERVICIO: UsuarioService
//
// Administra todas las operaciones relacionadas con los
// usuarios del sistema.
//
// Responsabilidades:
//
// • Registrar usuarios.
// • Buscar usuarios.
// • Actualizar información.
// • Eliminar usuarios.
// • Listar usuarios.
//
//==========================================================

type UsuarioService struct {

	usuarios map[int]*models.Usuario

}

//==========================================================
// Constructor
//==========================================================

func NuevoUsuarioService() *UsuarioService {

	return &UsuarioService{

		usuarios: make(map[int]*models.Usuario),

	}

}

//==========================================================
// Crear
//
// Registra un nuevo usuario.
//
//==========================================================

func (s *UsuarioService) Crear(

	usuario *models.Usuario,

) error {

	if _, existe := s.usuarios[usuario.GetID()]; existe {

		return errors.New("el usuario ya existe")

	}

	s.usuarios[usuario.GetID()] = usuario

	return nil

}

//==========================================================
// Buscar
//==========================================================

func (s *UsuarioService) Buscar(

	id int,

) (*models.Usuario, error) {

	usuario, existe := s.usuarios[id]

	if !existe {

		return nil, errors.New("usuario no encontrado")

	}

	return usuario, nil

}

//==========================================================
// Actualizar
//==========================================================

func (s *UsuarioService) Actualizar(

	usuario *models.Usuario,

) error {

	if _, existe := s.usuarios[usuario.GetID()]; !existe {

		return errors.New("usuario inexistente")

	}

	s.usuarios[usuario.GetID()] = usuario

	return nil

}

//==========================================================
// Eliminar
//==========================================================

func (s *UsuarioService) Eliminar(

	id int,

) error {

	if _, existe := s.usuarios[id]; !existe {

		return errors.New("usuario inexistente")

	}

	delete(s.usuarios, id)

	return nil

}

//==========================================================
// Listar
//==========================================================

func (s *UsuarioService) Listar() []*models.Usuario {

	lista := []*models.Usuario{}

	for _, usuario := range s.usuarios {

		lista = append(lista, usuario)

	}

	return lista

}