package application

import "publisher/src/domain"

type ObtenerPedido struct {
	repo domain.Ipedido
}

func NuevoObtenerPedido(repo domain.Ipedido) *ObtenerPedido {
	return &ObtenerPedido{repo: repo}
}

func (c *ObtenerPedido) Execute(id int) (*domain.Pedido, error) {
	return c.repo.GetByID(id)
}

