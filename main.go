package main

import (
	"log"
	"net/http"
	"os"

	"hacktiv8.com/assignment-3/handlers"
)

func main() {
	l := log.New(os.Stdout, "home-page", log.LstdFlags)

	sm := http.NewServeMux()
	hh := handlers.NewHome(l)
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:     ":8080",
		Handler:  sm,
		ErrorLog: l,
	}

	log.Fatal(s.ListenAndServe())
}
