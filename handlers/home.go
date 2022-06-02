package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"text/template"
)

type Home struct {
	l *log.Logger
}

type Element struct {
	Water int
	Wind  int
}

type data struct {
	Status Element
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (h *Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(rw, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	d := data{
		Status: Element{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		},
	}

	file, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile("data.json", file, 0644)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f := []string{
		"./templates/home.html",
	}

	ts, err := template.ParseFiles(f...)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(rw, d)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
