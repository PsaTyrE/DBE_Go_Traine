package controler

import (
	"myApp/helper"
	"myApp/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error {
	pass := c.Param("password")

	hash, _ := helper.HashPassword((pass))

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := model.CheckLogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(time.Hour * 72))

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
