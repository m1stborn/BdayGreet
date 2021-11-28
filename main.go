package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/m1stborn/BdayGreet/internal/pkg/handler"
	"github.com/m1stborn/BdayGreet/internal/pkg/model"

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
	model.MongoDB.Init(dbUri)

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/api/bdaygreet", HandleBdayGreet)

	log.Fatal(http.ListenAndServe(port, router))
}
