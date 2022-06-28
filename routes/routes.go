package routes

import (
	"net/http"

	"github.com/Haffif/web-echo/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/mahasiswa", controllers.FetchAllMahasiswa)
	e.POST("/inputMahasiswa", controllers.StoreMahasiswa)

	return e
}
