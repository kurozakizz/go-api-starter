package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	apiVersion = os.Getenv("API_VERION")
)

func getAPIPrefix() string {
	return "/v" + apiVersion
}

func main() {
	apiPrefix := getAPIPrefix()
	router := mux.NewRouter()
	router.HandleFunc(apiPrefix+"/pokemons", getPokemonListHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func getPokemonListHandler(w http.ResponseWriter, r *http.Request) {
	fruits := []string{"Pikachu", "Psyduck", "Happy"}
	json, _ := json.Marshal(fruits)
	responseJSON(w, r, json)
}

func responseJSON(w http.ResponseWriter, r *http.Request, json []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
