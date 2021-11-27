package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var (
	port  = ":" + os.Getenv("PORT")
	dbUri = os.Getenv("DB_URI")
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	fmt.Println(port, dbUri)

	router := httprouter.New()

	router.GET("/", Index)

	//log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(port, router))
}
