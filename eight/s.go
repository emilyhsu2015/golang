package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type food struct {
	Id     string
	Name   string
	Amount int
}

func main() {
	e := echo.New()
	e.GET("/amount/:id", getAmount)
	e.POST("/basic", basic)

	e.Logger.Fatal(e.Start(":1323"))
}

func getAmount(c echo.Context) error {
	test := []food{
		food{"f1", "巧克力餅乾", 80}, food{"f2", "魷魚絲", 180}}

	id := c.Param("id")
	var info string = ""

	for i := 0; i < len(test); i++ {
		if test[i].Id == id {
			info = test[i].Name + "," + strconv.Itoa(test[i].Amount)
		}
	}
	return c.String(http.StatusOK, info)
}

func basic(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "Name:"+name+", Email:"+email)
}
