package services

import (
	"e-commerce/models"
	"sync"
)	


//==========================================================
// SERVICIO: PedidoService
//==========================================================

type PedidoService struct {

	pedidos map[int]*models.Pedido
    
	siguienteID int

	mutex sync.RWMutex
}

//Constructor
func NuevoPedidoService() *PedidoService {

	return &PedidoService{

		pedidos: make(map[int]*models.Pedido),
        siguienteID: 1,
	}

}

//==========================================================
// Generador
//==========================================================
func (s *PedidoService) GenerarID() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.siguienteID

	s.siguienteID++

	return id
}

//==========================================================
// Crear
//==========================================================
func (s *PedidoService) Crear(

	pedido *models.Pedido,

) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.pedidos[pedido.GetID()] = pedido

}

//==========================================================
// Buscar
//==========================================================
func (s *PedidoService) Buscar(

	id int,

) *models.Pedido {
	s.mutex.RLock()
	defer s.mutex.RLock()

	return s.pedidos[id]

}

//==========================================================
// Listar
//==========================================================
func (s *PedidoService) Listar() []*models.Pedido {
	s.mutex.RLock()
	defer s.mutex.RLock()

	lista := []*models.Pedido{}

	for _, pedido := range s.pedidos {

		lista = append(lista, pedido)

	}

	return lista

}