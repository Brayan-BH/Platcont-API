package routes

import (
	"encoding/json"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/orm"

	"github.com/gorilla/mux"
)

func RutasAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {

	controller.SessionMgr.SetSessionVal(controller.SessionID, "login", true)

	w.Header().Set("Content-Type", "application-Json")
	response := controller.NewResponseManager()
	
	data_User := orm.NewQuerys("users").Where("email", "=", "email").Exec().All()

	if len(data_User) <= 0 {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contraseña Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["users"] = data_User
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
