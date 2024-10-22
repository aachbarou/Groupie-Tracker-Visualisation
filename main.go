package main

import (
	"fmt"
	"net/http"
	"time"

	Handler "groupie-tracker/Handlers"
)

func main() {
	go Handler.Startup()
	time.Sleep(8 * time.Second)
	http.HandleFunc("/", Handler.HAndleHOme)
	http.HandleFunc("/search", Handler.Handle_SearchBar)
	http.HandleFunc("/detailes", Handler.Handle_Detailes)
	http.HandleFunc("/filter", Handler.Handle_Filters)
	http.Handle("/Style.css" , http.FileServer(http.Dir("./Styles/")))
	http.Handle("/detailes.css" , http.FileServer(http.Dir("./Styles/")))
	fmt.Println("server  start at  Port 8080 : http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}