package controllers

import (
	"fmt"
	"net/http"

	"github.com/djarum76-bot/crud_post/models"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	claim := GetNilaiToken(c)
	id := fmt.Sprintf("%d", claim.Id)
	result, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
