package infraestructure

import (
	"publisher/core"
	"publisher/src/domain"
	"database/sql"
	"log"
	"fmt"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) RealizarPedido(p *domain.Pedido) error {
	query := "INSERT INTO cuenta (nombre, precio) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Precio)
	return err
}

func (r *MySQLRepository) GetByID(id int) (*domain.Pedido, error) {
	query := "SELECT id, nombre, precio FROM cuenta WHERE id = ?"
	row := r.conn.DB.QueryRow(query, id)

	var pedido domain.Pedido
	err := row.Scan(&pedido.Id, &pedido.Nombre, &pedido.Precio)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pedido con ID %d no encontrado", id)
		}
		log.Println("Error al obtener el pedido:", err)
		return nil, err
	}

	return &pedido, nil
}
