package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/database/orm"

	"github.com/gorilla/mux"
)

func RutasClientes(r *mux.Router) {
	s := r.PathPrefix("/clientes").Subrouter()
	s.Handle("/create", (http.HandlerFunc(insertCliente))).Methods("POST")
	s.Handle("/update/{uid}", (http.HandlerFunc(updateCliente))).Methods("PUT")


}


func insertCliente(w http.ResponseWriter, r *http.Request) {
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

	response.Data = _Clientes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	uid := params["uid"]
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
