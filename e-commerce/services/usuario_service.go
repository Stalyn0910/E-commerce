package services

import (
	"errors"
	"e-commerce/models"
	"sync"
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
	
	siguienteID int

	mutex sync.RWMutex
}

//==========================================================
// Constructor
//==========================================================

func NuevoUsuarioService() *UsuarioService {

	return &UsuarioService{

		usuarios:  make(map[int]*models.Usuario),
		siguienteID: 1,
	}

}

//==========================================================
// Crear
//
// Registra un nuevo usuario.
//
//==========================================================
func (s *UsuarioService) GenerarID() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.siguienteID

	s.siguienteID++

	return id
}

func (s *UsuarioService) Crear(

	usuario *models.Usuario,

) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	s.mutex.RLock()
	defer s.mutex.RUnlock()

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
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	lista := []*models.Usuario{}

	for _, usuario := range s.usuarios {

		lista = append(lista, usuario)

	}

	return lista

}