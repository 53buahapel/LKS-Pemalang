package web

import (
	"database/sql"
	"errors"
	"gogogo/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *model.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (username, password, role) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UserExists(username string) bool {
	stmt, err := r.db.Prepare("SELECT COUNT(*) FROM users WHERE username = ?")
	if err != nil {
		return false
	}

	var count int
	err = stmt.QueryRow(username).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	stmt, err := r.db.Prepare("SELECT id, username, password, role FROM users WHERE username = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user model.User
	err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *Repository) GetTodosByUserID(userID int) ([]model.Todo, error) {
	stmt, err := r.db.Prepare("SELECT t.id, t.title, t.done, u.id, u.username, u.password, u.role FROM todos t JOIN users u ON t.author_id = u.id WHERE u.id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.Author.ID, &todo.Author.Username, &todo.Author.Password, &todo.Author.Role)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *Repository) CreateTodo(todo *model.Todo) error {
	stmt, err := r.db.Prepare("INSERT INTO todos (title, done, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.Title, todo.Done, todo.Author.ID)
	if err != nil {
		return err
	}

	return nil
}
