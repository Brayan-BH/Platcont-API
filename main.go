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

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// router.HandleFunc("/static/img", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r)
	// 	file, err := ioutil.ReadFile("./public/code.png")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	w.Write(file)
	// })

	routes.RutasAuth(router)
	routes.RutasClientes(router)
	routes.RutasClientesProductos(router)
	routes.RutasFacturas(router)
	routes.RutasProductosDetalle(router)
	routes.RutaVersiones(router)

	fmt.Printf("server listening on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

func inits(w http.ResponseWriter, r *http.Request) {

	datos := map[string]string{"Api": "Platcont", "Version": "2.0.17", "Author": "Deybin Yoni Gil Perez"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datos)

}
