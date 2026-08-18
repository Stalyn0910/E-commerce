package services

import (
	"e-commerce/models"
	"sync"
)

//==========================================================
// SERVICIO: PagoService
//==========================================================

type PagoService struct {

	pagos map[int]*models.Pago
    
	siguienteID int
	mutex sync.RWMutex
}

//Constructor
func NuevoPagoService() *PagoService {

	return &PagoService{

		pagos: make(map[int]*models.Pago),
        siguienteID: 1,
	}

}
//==========================================================
// Generador
//==========================================================
func (s *PagoService) GenerarID() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.siguienteID

	s.siguienteID++

	return id
}

//==========================================================
// Registrar
//==========================================================
func (s *PagoService) Registrar(

	pago *models.Pago,

) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	pago.RegistrarPago()

	s.pagos[pago.GetID()] = pago

}

//==========================================================
// Buscar
//==========================================================
func (s *PagoService) Buscar(

	id int,

) *models.Pago {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.pagos[id]

}