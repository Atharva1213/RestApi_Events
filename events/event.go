package events

import (
	"Server_main/database"
	"errors"
)
type Event struct {
	ID     int64  `json:"id"`
	Name   string `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	UserId int64 `json:"userid"`
	Price int64 `json:"price" binding:"required"`
	TotalCount int64 `json:"total_count"`
}

func (event Event) Save(user_id int64) error {
	query := `INSERT INTO events (name,description,user_id,price) VALUES (?,?,?,?);`
	stm, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(event.Name, event.Description,user_id,event.Price)
	if err != nil {
		return err
	}
	return nil
}
func GetAllEvents() ([]Event, error) {
	// Prepare the SQL query
	query := `SELECT * FROM events;`
	result, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var events []Event

	for result.Next() {
		var event Event
		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.UserId,&event.Price,&event.TotalCount)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if len(events) == 0 {
		return []Event{}, nil
	}
	return events, nil
}
func GetEventByID(id int64) (Event, error) {
	// Prepare the SQL query
	query := `SELECT * FROM events WHERE id = ?`
	result, err := database.DB.Query(query, id)
	if err != nil {
		return Event{}, err
	}
	defer result.Close()

	var event Event

	if result.Next() {
		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.UserId,&event.Price,&event.TotalCount)
		if err != nil {
			return Event{}, err
		}
	} else {
		return Event{}, errors.New("Event not found")
	}

	return event, nil
}
func DeleteEvent(id int64) error {
	// Prepare the SQL query
	query := "DELETE FROM events WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (event Event) UpdatedEvent() error {
	query := `UPDATE events set name = ? ,description = ? ,user_id = ?, price = ? WHERE id = ?;`
	stm, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(event.Name, event.Description, event.UserId,event.Price,event.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdatedEventCount(id int64) error {
	query := `UPDATE events SET total_count = total_count + 1 WHERE id = ?;`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

