package services

import "p/models"

//==========================================================
// SERVICIO: PedidoService
//==========================================================

type PedidoService struct {

	pedidos map[int]*models.Pedido

}

func NuevoPedidoService() *PedidoService {

	return &PedidoService{

		pedidos: make(map[int]*models.Pedido),

	}

}

func (s *PedidoService) Crear(

	pedido *models.Pedido,

) {

	s.pedidos[pedido.GetID()] = pedido

}

func (s *PedidoService) Buscar(

	id int,

) *models.Pedido {

	return s.pedidos[id]

}

func (s *PedidoService) Listar() []*models.Pedido {

	lista := []*models.Pedido{}

	for _, pedido := range s.pedidos {

		lista = append(lista, pedido)

	}

	return lista

}