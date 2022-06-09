package models

import (
	"database/sql"
	"net/http"

	"github.com/djarum76-bot/crud_post/db"
)

func GetUser(id string) (Response, error) {
	var res Response
	var user User

	con := db.CreateCon()

	if _, err := con.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL NOT NULL PRIMARY KEY, username VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL);"); err != nil {
		return res, err
	}

	sqlStatement := "SELECT id, username FROM users WHERE id = ($1)"

	err := con.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Username)
	if err == sql.ErrNoRows {
		return res, err
	}
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil get data user"
	res.Data = user

	return res, nil
}
