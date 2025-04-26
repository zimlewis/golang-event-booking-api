package registration

import (
	"example.com/udemy_course/db"
)

type Registration struct {
	Id int64 `json:"id"`
	UserId int64 `json:"userId" bind:"required"`
	EventId int64 `json:"eventId" bind:"required"`
}

func Create(r Registration) error {
	query := `
	INSERT INTO Registration (UserId, EventId) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.UserId, r.EventId)

	return err
}

func Read(id int64) (*Registration, error) {
	return nil, nil
}

func ReadAll() ([]Registration, error) {
	return nil, nil
}

func Update(r Registration) error {
	return nil
}

func Delete(id int64) error {
	query := `
	DELETE FROM Registration WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
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

func New(uid, eid int64) *Registration {
	r := new(Registration)
	r.EventId = eid
	r.UserId = uid

	return r
}

func FindWithUserIdAndEventId(uid, eid int64) (*Registration, error) {
	query := `
	SELECT * FROM Registration WHERE UserId = ? AND EventId = ?;
	`
	var r Registration

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(uid, eid)
	err = row.Scan(&r.Id, &r.UserId, &r.EventId)

	return &r, err
}