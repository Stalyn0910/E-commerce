package services

import "p/models"

//==========================================================
// SERVICIO: PagoService
//==========================================================

type PagoService struct {

	pagos map[int]*models.Pago

}

func NuevoPagoService() *PagoService {

	return &PagoService{

		pagos: make(map[int]*models.Pago),

	}

}

func (s *PagoService) Registrar(

	pago *models.Pago,

) {

	pago.RegistrarPago()

	s.pagos[pago.GetID()] = pago

}

func (s *PagoService) Buscar(

	id int,

) *models.Pago {

	return s.pagos[id]

}