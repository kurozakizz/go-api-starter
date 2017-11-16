package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kurozakizz/go-api-starter/config"
	"github.com/kurozakizz/go-api-starter/middleware"
)

func main() {
	apiPrefix := config.GetAPIPrefix()
	mws := middleware.ChainMiddleware(middleware.WriteRequestLog)

	fmt.Println("Running API " + apiPrefix + " ...")
	http.HandleFunc(apiPrefix+"/pokemons", mws(getPokemonListHandler))
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
