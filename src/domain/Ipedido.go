package domain

type Ipedido interface{
	RealizarPedido(pedido *Pedido)error
}