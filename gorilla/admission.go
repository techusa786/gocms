package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/techusa786/nmhutil"
)

func main() {
	fmt.Println("Hello Gorilla Mux")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexHandler)

	// all persons handler
	router.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling /persons route")
		fmt.Fprintln(w, "calling /persons route")
	})

	// person handler with id
	router.HandleFunc("/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling /persons/{id} route")
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("id:", id)
		fmt.Fprintln(w, id)
	})

	// person handler with firstname/lastname
	router.HandleFunc("/person/{firstname}/{lastname}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling /person/{firstname}/{lastname} route")
		vars := mux.Vars(r)
		firstname := vars["firstname"]
		lastname := vars["lastname"]
		fmt.Println("firstname:", firstname, "  lastname:", lastname)
		fmt.Fprintln(w, firstname, lastname)
	})

	// person handler for /person/{id}/{firstname}/{lastname}
	router.HandleFunc("/person/{id}/{firstname}/{lastname}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling /person/{id}/{firstname}/{lastname} route")
		vars := mux.Vars(r)
		id := vars["id"]
		firstname := vars["firstname"]
		lastname := vars["lastname"]
		fmt.Println("id:", id, " firstname:", firstname, " lastname:", lastname)
		fmt.Fprintln(w, "handler for /person/{id}/{firstname}/{lastname}")
		fmt.Fprintln(w, id, firstname, lastname)
	})

	error := http.ListenAndServe(":3000", router)
	nmhutil.CheckError(error)
	fmt.Println("gorilla mux listening on port 3000")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling /Index route")
	fmt.Fprintln(w, "Index root");
}



//func main() {
//
//	router := mux.NewRouter().StrictSlash(true)
//	router.HandleFunc("/", Index)
//	router.HandleFunc("/todos", TodoIndex)
//	router.HandleFunc("/todos/{todoId}", TodoShow)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}
//
//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Welcome!")
//}
//
//func TodoIndex(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Todo Index!")
//}
//
//func TodoShow(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	todoId := vars["todoId"]
//	fmt.Fprintln(w, "Todo show:", todoId)
//}
