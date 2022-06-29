package controllers

import (
	"net/http"
	"strconv"

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

func UpdateMahasiswa(c echo.Context) error {

	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_telepon := c.FormValue("no_telepon")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	result, err := models.UpdateMahasiswa(conv_id, nama, alamat, no_telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)
}
