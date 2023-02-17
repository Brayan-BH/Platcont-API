package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/orm"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	// Get the request body
	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//objeto map
	body := make(map[string]interface{})
	
	json.Unmarshal(req_body, &body)

	dataUser := orm.NewQuerys("Seguridad").Select().Where("email;", "=", body["email"]).Exec().One()
	if len(dataUser) <= 0 {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contarseña Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//int8, int64, int32

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["password_admin"].(string)), []byte(body["password_admin"].(string)))
	if err != nil {
		response.Msg = "Usuario y Contraseña Inconrrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	controller.SessionMgr.SetSessionVal(controller.SessionID, "login", true)

	response.Data["users"] = dataUser
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
