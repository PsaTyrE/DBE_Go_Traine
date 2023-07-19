package route

import (
	"myApp/controler"
	"myApp/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})

	e.GET("/pegawai", controler.FetchAllPegawai, middleware.IsAuth)
	e.POST("/pegawai", controler.StorePegawai)
	e.PUT("/pegawai", controler.UpdatePegawai)
	e.DELETE("/pegawai", controler.DeletePegawai)

	e.GET("/generate-hash/:password", controler.GenerateHashPassword)
	e.POST("/login", controler.CheckLogin)

	return e
}
