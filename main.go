package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kurozakizz/go-api-starter/middleware"
)

var (
	apiPrefix = "/v" + os.Getenv("API_VERION")
)

func main() {
	fmt.Println("Running API " + apiPrefix + " ...")
	http.HandleFunc(apiPrefix+"/pokemons", middleware.WriteRequestLog(getPokemonListHandler))
	http.ListenAndServe(":5000", nil)
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
