package models

import "api_pgsql/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := `UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4`

	res, err := conn.Exec(query, todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
