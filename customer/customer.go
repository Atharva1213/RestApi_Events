package customer 

import (
	"Server_main/database"
)

type Customer struct {
	ID     int64  `json:"id"`
	Name   string `json:"name" binding:"required"`
	EventId int64 `json:"eventid" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	PaymentMode string   `json:"paymentmode" binding:"required"`
}

func (customer Customer) Register() error {
	query := `INSERT INTO Register (name,event_id,paymentmode) VALUES (?,?,?);`
	stm, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(customer.Name,customer.EventId,customer.PaymentMode)
	if err != nil {
		return err
	}
	return nil
}

func (customer Customer) CheckRegister() (bool, error) {
	query := `SELECT * FROM Register WHERE email = ? AND event_id = ?;`
	row := database.DB.QueryRow(query, customer.Email, customer.EventId)
	var result struct{}
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
