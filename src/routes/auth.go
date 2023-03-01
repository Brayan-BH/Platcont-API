package routes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/database/orm"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RutasAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
	s.HandleFunc("/logout", logout).Methods("PUT")
	s.HandleFunc("/first-step", registerFirst).Methods("POST")
	s.HandleFunc("/second-step", registerSecond).Methods("POST")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func logout(w http.ResponseWriter, r *http.Request) {
	//TODO: Logout
	controller.SessionMgr.EndSessionBy(controller.SessionID)

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	if close == true {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
	return

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	controller.SessionID = controller.SessionMgr.StartSession(w, r)

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

	dataUser := orm.NewQuerys("users").Select().Where("email", "=", body["email"]).Exec().One()
	if len(dataUser) <= 0 {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contraseña Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["password_admin"].(string)), []byte(body["password_admin"].(string)))
	if err != nil {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	// Establecer el valor "login" en la sesión
	controller.SessionMgr.SetSessionVal(controller.SessionID, "login", true)
	controller.SessionMgr.SetSessionVal(controller.SessionID, "id_user", dataUser["id_user"].(string))

	response.Data["users"] = dataUser
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func registerFirst(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Users_GetSchema()
	_Clientes := orm.SqlExec{}
	err = _Clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Clientes.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Clientes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func registerSecond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clients_GetSchema()
	_Clientes := orm.SqlExec{}
	err = _Clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Clientes.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	returnData := _Clientes.Data[0]
	response.Data = returnData
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
