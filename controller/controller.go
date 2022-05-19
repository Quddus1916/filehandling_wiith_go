package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func Sendfile(c echo.Context) error {
	msg := "send file working"
	return c.JSON(http.StatusOK, msg)
}

func Getfiles(c echo.Context) error {
	msg := "get file working"
	return c.JSON(http.StatusOK, msg)
}
