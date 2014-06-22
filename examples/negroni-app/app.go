package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/meatballhat/go-newrelic/negroni/nr"
)

var (
	addr = ":" + os.Getenv("PORT")
)

func main() {
	n := negroni.Classic()
	err := nr.Configure(n)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(map[string]string{"whatever": "hosers"})
		fmt.Fprintf(w, string(j))
	}).Methods("GET")

	if addr == ":" {
		addr = ":3000"
	}

	n.UseHandler(r)
	n.Run(addr)
}
