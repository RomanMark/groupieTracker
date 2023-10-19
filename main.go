package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8060"

func main() {

	fmt.Println("Starting server at port", port)
	fmt.Println("http://localhost:8060/")
	getData()

	http.HandleFunc("/", home)

	http.HandleFunc("/artist/", artistPage)
	log.Fatal(http.ListenAndServe(":8060", nil))

}
