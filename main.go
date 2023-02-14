package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"platcont/src/controller"
	"platcont/src/helper"
	"platcont/src/middleware"
	"platcont/src/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	controller.SessionMgr = helper.NewSessionMgr("cookie-token", 3600)
	router := mux.NewRouter().StrictSlash(true)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	middleware.EnableCORS(router)
	router.HandleFunc("/", inits)
	routes.RutasAuth(router)
	routes.RutasClientes(router)

	fmt.Printf("server listening on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func inits(w http.ResponseWriter, r *http.Request) {

	datos := map[string]string{"Api": "Platcont", "Version": "2.0.17", "Author": "Deybin Yoni Gil Perez"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datos)

}
