package domain

type Ipedido interface{
	RealizarPedido(pedido *Pedido)error
	// VerPedido()([]Pedido,error)
	// BorrarPedido(id int)error
	// ActualizarPedido(id int,pedido *Pedido)error
}