package model

import "gogogo/database"

type Todo struct {
	ID     int
	Title  string
	Done   bool
	Author User
}

func (t *Todo) MarkDone() error {
	stmt, err := database.DB.Prepare("UPDATE todos SET done = true WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.ID)
	if err != nil {
		return err
	}

	return nil
}
