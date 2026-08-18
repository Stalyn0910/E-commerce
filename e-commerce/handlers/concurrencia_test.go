package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"e-commerce/services"
)

func TestConcurrenciaUsuarios(t *testing.T) {

	service := services.NuevoUsuarioService()

	handler := NuevoUsuarioHandler(service)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {

		wg.Add(1)

		go func(i int) {

			defer wg.Done()

			request := httptest.NewRequest(
				http.MethodGet,
				"/usuarios",
				nil,
			)

			recorder := httptest.NewRecorder()

			handler.Listar(
				recorder,
				request,
			)

			if recorder.Code != http.StatusOK {

				t.Errorf(
					"goroutine %d: se esperaba 200, se obtuvo %d",
					i,
					recorder.Code,
				)
			}
		}(i)
	}
	wg.Wait()
}