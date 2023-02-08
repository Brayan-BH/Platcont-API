package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"platcont/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomeHandler)
	//Rutas de autentificacion
	routes.RutasAuth(r)
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplicattion/json")
	data := map[string]interface{}{"api": "platcontApi", "version": 1.1}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
