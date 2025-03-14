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

// func (r *MySQLRepository) Delete(p int)error{
// 	id:=p
// 	query := "DELETE FROM Users WHERE id = ?"
// 	_,err :=r.conn.DB.Exec(query,id)
// 	return err
// }

// func (r *MySQLRepository) Update(id int,p *domain.Cuenta)error{
// 	query := "UPDATE cuenta SET titular = ?, saldo = ?, moneda=? WHERE id = ?"
//     _, err := r.conn.DB.Exec(query,p.Titular,p.Saldo,p.Moneda,id)
// 	return err
// }

// func (r *MySQLRepository) GetAll() ([]domain.Cuenta, error) {
// 	query := "SELECT titular, saldo, moneda, creado_en FROM cuenta"
// 	rows, err := r.conn.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
    
// 	var cuentas []domain.Pedido
// 	for rows.Next() {
// 		var cuenta domain.Pedido
// 		if err := rows.Scan(&cuenta.Titular, &cuenta.Saldo,&cuenta.Moneda,&cuenta.Creado_en); err != nil {
// 			return nil, err
// 		}
// 		cuentas = append(cuentas, cuenta)  
// 	}

// 	return cuentas, nil
// }


// func (r *MySQLRepository) Deposit(id int,mount float64)error{
// 	query := "UPDATE cuenta SET saldo=saldo + ? where id=?"
// 	_, err := r.conn.DB.Query(query,mount,id)
// 	return err
// }
// func (r *MySQLRepository) Transfer(fromId int, toId int, mount float64) error {
  
//     tx, err := r.conn.DB.Begin()
//     if err != nil {
//         fmt.Println("❌ Error iniciando transacción:", err)
//         return err
//     }

//     var saldoActual float64

//     err1 := tx.QueryRow("SELECT saldo FROM cuenta WHERE id=?", fromId).Scan(&saldoActual)
//     if err1 != nil {
//         fmt.Println("❌ Error: No se encontró el usuario con ID:", fromId)
//         tx.Rollback()
//         return errors.New("no se encontró al usuario")
//     }

    
//     if saldoActual < mount {
//         fmt.Println("❌ Error: Saldo insuficiente. Saldo actual:", saldoActual, "Monto requerido:", mount)
//         tx.Rollback()
//         return errors.New("no tienes dinero suficiente")
//     }

 
//     _, err = tx.Exec("UPDATE cuenta SET saldo = saldo - ? WHERE id = ?", mount, fromId)
//     if err != nil {
//         fmt.Println("❌ Error al descontar saldo de la cuenta", fromId, ":", err)
//         tx.Rollback()
//         return err
//     }

    
//     var exists int
//     err2 := tx.QueryRow("SELECT COUNT(*) FROM cuenta WHERE id=?", toId).Scan(&exists)
//     if err2 != nil || exists == 0 {
//         fmt.Println("❌ Error: El usuario destino con ID", toId, "no existe")
//         tx.Rollback()
//         return errors.New("el usuario destino no existe")
//     }

    
//     _, err = tx.Exec("UPDATE cuenta SET saldo = saldo + ? WHERE id = ?", mount, toId)
//     if err != nil {
//         fmt.Println("❌ Error al agregar saldo a la cuenta", toId, ":", err)
//         tx.Rollback()
//         return err
//     }

  
//     err = tx.Commit()
//     if err != nil {
//         fmt.Println("❌ Error al confirmar la transacción:", err)
//         return err
//     }

//     fmt.Println("✅ Transferencia exitosa de", mount, "de", fromId, "a", toId)
//     return nil
// }
