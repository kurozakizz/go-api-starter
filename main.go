package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var (
	apiVersion = os.Getenv("API_VERION")
)

func getAPIPrefix() string {
	return "/v" + apiVersion
}

func main() {
	apiPrefix := getAPIPrefix()
	http.HandleFunc(apiPrefix+"/pokemons", requestLogger(getPokemonListHandler))
	http.ListenAndServe(":5000", nil)
}

func requestLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)
		h(w, r)
	}
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
