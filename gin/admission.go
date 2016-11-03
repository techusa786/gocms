package main

import (
	//"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/techusa786/nmhutil"
	"net/http"
)

type Person struct {
	Id         int
	First_Name string
	Last_Name  string
}

func main() {
	// connect to db and ping it
	db, err := sql.Open("mysql", "root:wizard786@tcp(127.0.0.1:3306)/gotest")
	nmhutil.CheckError(err)

	defer db.Close()

	// let us check if connection is available
	err = db.Ping()
	nmhutil.CheckError(err)

	// gin router
	router := gin.Default()
	fmt.Println(router)

	router.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"language": "golang",
		//	"framework": "gin",
		//	"message": "Hello github.com/gin-gonic/gin!",
		//})
		c.String(http.StatusOK, "Hello %s", "github.com/gin-gonic/gin")
	})

	// route to get all persons in db
	router.GET("/persons", func(c* gin.Context) {
		var person Person
		var persons []Person

		query := "select id, first_name, last_name from person"
		rows, err := db.Query(query)
		nmhutil.CheckError(err)
		defer rows.Close()

		for rows.Next() { //? this for loop
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
			nmhutil.CheckError(err)
			persons = append(persons, person)
		}
		// send json
		c.JSON(http.StatusOK, gin.H{ //? check the syntax
			"result": persons,
			"count": len(persons),
		})
	})


	// GET a specific person detail with multiple params
	//router.GET("/person/:firstname/lastname", func(c* gin.Context) {
	//	var person Person
	//	var result gin.H
	//	firstName := c.Param("firstname")
	//	lastName := c.Param("lastname")
	//	query := "select id, first_name, last_name from person where first_name = ? and last_name = ?"
	//	row := db.QueryRow(query, firstName, lastName)
	//	err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
	//	nmhutil.CheckError(err)
	//
	//	if err != nil {
	//		result = gin.H{
	//			"result": nil,
	//			"count": 0,
	//		}
	//	} else {
	//		result = gin.H{
	//			"result": person,
	//			"count": 1,
	//		}
	//	}
	//	c.JSON(http.StatusOK, result)
	//})

	// GET a specific person detail
	router.GET("/person/:id", func(c* gin.Context) {
		var person Person
		var result gin.H
		id := c.Param("id")
		query := "select id, first_name, last_name from person where id = ?"
		row := db.QueryRow(query, id)
		err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		nmhutil.CheckError(err)

		if err != nil {
			result = gin.H{
				"result": nil,
				"count": 0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count": 1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// start gin
	router.Run(":3000")

	// insert into person (first_name, last_name) values(?,?)
}
