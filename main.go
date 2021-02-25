package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	servergo "github.com/francesc2509/http-wrapper"
)

type Item struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

var items []*Item = []*Item{}

func main() {
	r := servergo.New()

	for id := 1; id <= 10; id++ {
		name := fmt.Sprintf("Item %d", id)

		items = append(items, &Item{Name: name, Id: id})
	}

	r.HandleFunc("/", get).Methods("OPTIONS", "GET")
	r.HandleFunc("/add", add).Methods("OPTIONS", "POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(items)
}

func add(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	item := &Item{}
	json.Unmarshal(bytes, item)

	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(item)

	defer r.Body.Close()
}
