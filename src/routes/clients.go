package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/database/orm"
	"platcont/src/libraries/library"
	"platcont/src/middleware"

	"github.com/gorilla/mux"
)

func RutasClientes(r *mux.Router) {
	s := r.PathPrefix("/clientes").Subrouter()
	s.Handle("/info", middleware.Autentication(http.HandlerFunc(GetOneClient))).Methods("GET")
	// s.Handle("/register", (http.HandlerFunc(RegisterCliente))).Methods("POST")
	// s.Handle("/create", (http.HandlerFunc(CreateUser))).Methods("POST")
	s.Handle("/update", (http.HandlerFunc(UpdateCliente))).Methods("PUT")
}

func GetOneClient(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	id_clie := library.GetSession_key("id_user")

	//get allData from database
	dataUser := orm.NewQuerys("clients").Select().Where("id_clie", "=", id_clie).Exec().One()

	if len(dataUser) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["info"] = dataUser

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	uid := library.GetSession_key("id_user")
	if uid == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"uid": uid}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Clients_GetSchema()
	_Clientes := orm.SqlExec{}
	err = _Clientes.New(data_update, table).Update(schema)
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
