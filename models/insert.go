package models

import "api_pgsql/db"

func Insert(todo Todo) (id int64, err error) {
	db, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(query, todo.Title, todo.Description, todo.Done).Scan(&id)
	if err != nil {
		return 0, err
	}

	return
}
