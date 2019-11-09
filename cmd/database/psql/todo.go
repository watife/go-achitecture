package psql

import (
	"database/sql"
	"errors"
	"fakorede-bolu/go-ach/cmd/todo"
	"fmt"
	"log"

	// "log"

	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func NewPostgresTodoRepository(db *sql.DB) todo.Repository {
	return &repository{
		db,
	}
}

func (r *repository) Create(t *todo.Model) (*todo.Model, error) {
	stmt := `INSERT INTO todos (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	todo := &todo.Model{}

	row := r.db.QueryRow(stmt,
		t.Title, t.Description, t.Status)

	fmt.Println(row)

	err := row.Scan(&todo.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err

	}
	return todo, nil
}

func (r *repository) FindByID(id string) (*todo.Model, error) {
	todo := new(todo.Model)
	err := r.db.QueryRow("SELECT id, title, description, status FROM todos where id=$1", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)
	if err != nil {
		panic(err)
	}
	return todo, nil
}

func (r *repository) FindAll() (t []*todo.Model, err error) {
	stmt := `SELECT * FROM todos;`
	rows, err := r.db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*todo.Model{}

	for rows.Next() {

		t := &todo.Model{}

		if err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status); err != nil {
			log.Print(err)
			return nil, err
		}

		todos = append(todos, t)

	}
	return todos, nil
}
