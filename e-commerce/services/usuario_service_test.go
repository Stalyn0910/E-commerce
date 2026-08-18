package services

import (
	"sync"
	"testing"

	"e-commerce/models"
)
func TestUsuarioServiceCrear(t *testing.T) {

	service := NuevoUsuarioService()

	usuario := models.NuevoUsuario(
		service.GenerarID(),
		"Stalyn",
		"stalyn@gmail.com",
		"123456",
		"Cliente",
	)

	err := service.Crear(usuario)

	if err != nil {
		t.Fatalf(
			"no se pudo crear el usuario: %v",
			err,
		)
	}

	resultado, err := service.Buscar(usuario.GetID())

	if err != nil {
		t.Fatalf(
			"no se encontró el usuario creado: %v",
			err,
		)
	}

	if resultado.GetNombre() != "Stalyn" {

		t.Errorf(
			"se esperaba el nombre Stalyn, se obtuvo %s",
			resultado.GetNombre(),
		)
	}
}
func TestUsuarioServiceDuplicado(t *testing.T) {

	service := NuevoUsuarioService()

	id := service.GenerarID()

	usuario1 := models.NuevoUsuario(
		id,
		"Usuario 1",
		"usuario1@gmail.com",
		"123456",
		"Cliente",
	)

	usuario2 := models.NuevoUsuario(
		id,
		"Usuario 2",
		"usuario2@gmail.com",
		"123456",
		"Cliente",
	)

	if err := service.Crear(usuario1); err != nil {

		t.Fatalf(
			"no se pudo crear el primer usuario: %v",
			err,
		)
	}

	if err := service.Crear(usuario2); err == nil {

		t.Error(
			"se permitió crear un usuario con ID duplicado",
		)
	}
}
func TestUsuariosConcurrentes(t *testing.T) {

	service := NuevoUsuarioService()

	const cantidadUsuarios = 100

	var wg sync.WaitGroup

	wg.Add(cantidadUsuarios)

	for i := 0; i < cantidadUsuarios; i++ {

		go func() {

			defer wg.Done()

			id := service.GenerarID()

			usuario := models.NuevoUsuario(
				id,
				"Usuario Concurrente",
				"usuario@gmail.com",
				"123456",
				"Cliente",
			)

			if err := service.Crear(usuario); err != nil {

				t.Errorf(
					"error creando usuario concurrente: %v",
					err,
				)
			}

		}()

	}

	wg.Wait()

	usuarios := service.Listar()

	if len(usuarios) != cantidadUsuarios {

		t.Errorf(
			"se esperaban %d usuarios, se obtuvieron %d",
			cantidadUsuarios,
			len(usuarios),
		)
	}
}
