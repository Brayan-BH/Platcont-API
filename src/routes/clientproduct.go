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

func RutasClientesProductos(r *mux.Router) {
	s := r.PathPrefix("/productos").Subrouter()
	s.Handle("/list", (http.HandlerFunc(allProducto))).Methods("GET")
	s.Handle("/info/{id_clie}", (http.HandlerFunc(oneProduct))).Methods("GET")
	s.Handle("/create", (http.HandlerFunc(insertProduct))).Methods("POST")
	s.Handle("/update/{id_clie}", (http.HandlerFunc(updateProduct))).Methods("PUT")

}

func allProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_client_product := orm.NewQuerys("Clientproducts").Select().Exec().All()

	if len(data_client_product) <= 0 {
		controller.ErrorsSuccess(w, errors.New("No se encontraron resultados para la consulta"))
		return
	}

	response.Data["clientProducts"] = data_client_product
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clientproducts_GetSchema()
	_Client_Products := orm.SqlExec{}
	err = _Client_Products.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Client_Products.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Client_Products.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_clie := params["id_clie"]
	if id_clie == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"id_clie": id_clie}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Clientproducts_GetSchema()
	_Client_Products := orm.SqlExec{}
	err = _Client_Products.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Client_Products.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Client_Products.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func oneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_clie := params["id_clie"]
	if id_clie == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	//get allData from database
	data_cliente := orm.NewQuerys("ClientProducts").Select("multi, users, date_facture, modulos, host").Where("id_clie", "=", id_clie).Exec().One()

	if len(data_cliente) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data = data_cliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
