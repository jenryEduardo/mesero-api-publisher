package domain

type Ipedido interface{
	RealizarPedido(pedido *Pedido)error
	GetByID(id int) (*Pedido, error)

}