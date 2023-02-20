package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"platcont/src/controller"
	"platcont/src/database/orm"

	"github.com/gorilla/mux"
)

func RutasClientesProductos(r *mux.Router) {
	s := r.PathPrefix("/productos").Subrouter()
	s.Handle("/list", (http.HandlerFunc(allProducto))).Methods("GET")
	s.Handle("/info/{id_clipd}", (http.HandlerFunc(oneProduct))).Methods("GET")
	// s.Handle("/create", (http.HandlerFunc(insertProduct))).Methods("POST")
	s.Handle("/update/{id_clipd}", (http.HandlerFunc(updateCliente))).Methods("PUT")

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

func oneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_clipd := params["id_clipd"]
	if id_clipd == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	//get allData from database
	data_cliente := orm.NewQuerys("Clients").Select("multi, users, date_facture").Where("id_clipd", "=", id_clipd).Exec().One()

	if len(data_cliente) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data = data_cliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
