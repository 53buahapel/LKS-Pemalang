package model

import (
	"gogogo/database"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

func (u *User) ChangeRole(role string) error {
	if u.Role == role {
		return nil
	}
	stmt, err := database.DB.Prepare("UPDATE users SET role = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(role, u.ID)
	if err != nil {
		return err
	}

	return nil
}
