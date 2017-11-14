package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/pokemons", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fruits := []string{"Pikachu", "Psyduck", "Happy"}
		json, _ := json.Marshal(fruits)

		w.Write(json)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
