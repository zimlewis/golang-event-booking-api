package event

import (
	"fmt"
	"time"

	"example.com/udemy_course/db"
)

type Event struct {
	Id int64 `json:"id"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location string `json:"location" binding:"required"`
	DateTime time.Time `json:"datetime" binding:"required"`
	UserId int64 `json:"userId"`
}

func Create(e Event) error {
	query := `
	INSERT INTO Event (Name, Description, Location, DateTime, UserId) VALUES
	(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	return err
}

func ReadAll() ([]Event, error) {
	query := `
	SELECT * FROM Event
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event = []Event{}

	for rows.Next() {
		var e Event
		rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		events = append(events, e)
	}

	return events, nil
}

func Read(id int64) (*Event, error) {
	query := `
	SELECT * FROM Event
	WHERE Id = ?
	`
	row := db.DB.QueryRow(query, id)

	fmt.Println(row)

	var e Event
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		return nil, err
	}
	

	return &e, nil
}

func Update(e Event) error {
	query := `
	UPDATE Event 
	SET Name = ?, Description = ?, Location = ?, DateTime = ?, UserId = ?
	WHERE Id = ?
	`


	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.Id)

	return err
}

func Delete(id int64) error {
	query := `DELETE FROM Event WHERE Id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func New(id int64, name, description, location string, userId int64) *Event {
	event := new(Event)

	event.Id = id
	event.Name = name
	event.Description = description
	event.Location = location
	event.UserId = userId

	return event
}

