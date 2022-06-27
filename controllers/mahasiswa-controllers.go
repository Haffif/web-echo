package controllers

import (
	"net/http"

	"github.com/Haffif/web-echo/models"
	"github.com/labstack/echo"
)

func FetchAllMahasiswa(c echo.Context) error {
	result, err := models.FetchAllMahasiswa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
