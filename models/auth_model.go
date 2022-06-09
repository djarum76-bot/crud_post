package models

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/djarum76-bot/crud_post/db"
	"github.com/djarum76-bot/crud_post/helpers"

	"github.com/golang-jwt/jwt"
)

func Register(username string, password string) (ResponseToken, bool, error) {
	var user User
	var err error
	var res ResponseToken
	lastInsertID := 0

	con := db.CreateCon()

	if _, err := con.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL NOT NULL PRIMARY KEY, username VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL);"); err != nil {
		return res, false, err
	}

	sqlStatement := `INSERT INTO users (username,password) VALUES ($1,$2) RETURNING id`

	err = con.QueryRow(sqlStatement, username, password).Scan(&lastInsertID)
	if err != nil {
		return res, false, err
	}

	// stmt, err := con.Prepare(sqlStatement)
	// if err != nil {
	// 	return res, false, err
	// }

	// _, err = stmt.Exec(username, password)
	// if err != nil {
	// 	return res, false, err
	// }

	user.Id = lastInsertID
	user.Username = username

	claims := &JwtCustomClaims{
		user.Id,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, false, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Register"
	res.Data = user
	res.Token = t

	return res, true, nil
}

func Login(username string, password string) (ResponseToken, bool, error) {
	var user User
	var pwdHash string
	var res ResponseToken

	con := db.CreateCon()

	if _, err := con.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL NOT NULL PRIMARY KEY, username VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL);"); err != nil {
		return res, false, err
	}

	sqlStatement := "SELECT * FROM users WHERE username = ($1)"

	err := con.QueryRow(sqlStatement, username).Scan(&user.Id, &user.Username, &pwdHash)
	if err == sql.ErrNoRows {
		return res, false, err
	}
	if err != nil {
		return res, false, err
	}

	match, err := helpers.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, false, err
	}

	claims := &JwtCustomClaims{
		user.Id,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, false, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Login"
	res.Data = user
	res.Token = t

	return res, true, nil
}
