package infraestructure


import (

	"publisher/core"
	"publisher/src/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}


func NewMySQLRepository() *MySQLRepository {
	conn := core. GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) RealizarPedido(p *domain.Pedido) error {
	query := "INSERT INTO cuenta (nombre,precio) VALUES (?,?)"
	_, err := r.conn.DB.Exec(query, p.Nombre,p.Precio)
	return err
}
