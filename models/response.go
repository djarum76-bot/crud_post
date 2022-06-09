package models

import "github.com/golang-jwt/jwt"

type ResponseToken struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

type Response struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
