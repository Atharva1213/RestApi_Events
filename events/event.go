package events

import (
	"Server_main/database"
	"errors"
)

var Events = []Event{}

type Event struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Price  int64  `json:"price" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (event Event) Save() error {
	// Prepare the SQL query
	query := `INSERT INTO events (name, price, author) VALUES (?, ?, ?);`
	stm, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(event.Name, event.Price, event.Author)
	if err != nil {
		return err
	}
	return nil
}
func GetAllEvents() ([]Event, error) {
	// Prepare the SQL query
	query := `SELECT * FROM events`
	result, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var events []Event

	for result.Next() {
		var event Event
		err := result.Scan(&event.ID, &event.Name, &event.Price, &event.Author)
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
		err := result.Scan(&event.ID, &event.Name, &event.Price, &event.Author)
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

