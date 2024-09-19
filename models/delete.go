package models

import "api_pgsql/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := `DELETE FROM todos WHERE id = $1`

	res, err := conn.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
