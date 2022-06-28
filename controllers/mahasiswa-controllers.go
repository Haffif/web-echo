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

func StoreMahasiswa(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_telepon := c.FormValue("no_telepon")

	result, err := models.StoreMahasiswa(nama, alamat, no_telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)
}
