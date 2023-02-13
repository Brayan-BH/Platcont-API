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
	// s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")

}

func login(w http.ResponseWriter, r *http.Request) {

	controller.SessionMgr.SetSessionVal(controller.SessionID, "login", true)

	w.Header().Set("Content-Type", "application-Json")
	response := controller.NewResponseManager()
	
	data_User := orm.NewQuerys("users").Select().Where("email", "=", "email")

	if len(data_User.Query) <= 0 {
		response.Msg = "Usuario y Contraseña Inconrrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contraseña Inconrrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["clientes"] = data_User
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
