package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"fmt"
	"time"
	"github.com/labstack/echo/middleware"
)

//type (
//	user struct {
//		ID   int    `json:"id"`
//		Name string `json:"name"`
//	}
//)
//
//var (
//	users = map[int]*user{}
//	seq   = 1
//)

type Person struct {
	Id int `json:"id" xml:"id" form:"id"`
	FirstName string `json:"firstname" xml:"firstname" form:"firstname"`
	LastName  string `json:"lastname" xml:"lastname" form:"lastname"`
}


func main() {
	e := echo.New()

	e.Get("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello Echo!")
	})

	// person route
	e.Get("/persons", func(ctx echo.Context) error {
		person := new(Person)
		person.Id = 111
		person.FirstName = "Naveed"
		person.LastName = "Anwar"

		if err := ctx.Bind(person); err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, person)
		//return ctx.XML(http.StatusCreated, person)

	})

	// person route
	e.Get("/persons/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		return ctx.String(http.StatusOK, "id:" + id)
	})

	// person route
	e.Get("/persons/:firstname/:lastname", func(ctx echo.Context) error {
		firstName := ctx.Param("firstname")
		lastName := ctx.Param("lastname")
		output := "route: /person/:firstname/:lastname \n"
		ctx.String(http.StatusOK, output)
		output = "firstname:" + firstName + "\t" + "lastname:" + lastName
		return ctx.String(http.StatusOK, output)
	})

	e.Get("/persons/:id/profile", func(ctx echo.Context) error {
		id := ctx.Param("id")
		output := "route: /persons/:id/profile \n"
		ctx.String(http.StatusOK, output)
		output = "id:" + id
		return ctx.String(http.StatusOK, output)
	})

	e.Get("/persons/:firstname/:lastname/profile", func(ctx echo.Context) error {
		firstName := ctx.Param("firstname")
		lastName := ctx.Param("lastname")
		output := "route: /person/:firstname/:lastname/profile \n"
		ctx.String(http.StatusOK, output)
		output = "firstname:" + firstName + "\t" + "lastname:" + lastName
		return ctx.String(http.StatusOK, output)
	})

	message := time.Now().Format(time.RFC850)
	fmt.Println("starting echo server at port 3000 on " + message)

	e.Use(middleware.Logger())
	e.Run(standard.New(":3000"))
}


//"/teachers" for the list
//"/teachers/:id" for an item in the list
//"/teachers/:id/profile
//github.com/
//github.com/issues
//github.com/:user
//github.com/:user/:repo
//
//twitter.com/mentions
//twitter.com/i/notifications
//twitter.com/:username
//
//www.facebook.com/events/upcoming
//www.facebook.com/messages/
//www.facebook.com/:username