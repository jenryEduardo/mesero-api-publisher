package application

import "publisher/src/domain"

type RealizarPedidos struct {
	repo domain.Ipedido
}


func NuevoPedido(repo domain.Ipedido)*RealizarPedidos{
	return &RealizarPedidos{repo: repo}
}


func (c *RealizarPedidos)Execute(pedido domain.Pedido)error{
	return c.repo.RealizarPedido(&pedido)
}