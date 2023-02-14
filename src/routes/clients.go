package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"platcont/src/controller"
	"platcont/src/database/orm"
	"platcont/src/middleware"

	"github.com/gorilla/mux"
)

func RutasClientes(r *mux.Router) {
	s := r.PathPrefix("/clientes").Subrouter()
	s.Handle("/list", middleware.Autentication(http.HandlerFunc(allCliente))).Methods("GET")
}

func allCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_cliente := orm.NewQuerys("requ_clientes").Select("c_docu,n_docu,l_clie").OrderBy("n_docu").Exec().All()

	if len(data_cliente) <= 0 {
		controller.ErrorsSuccess(w, errors.New("No se encontraron resultados para la consulta"))
		return
	}

	response.Data["clients"] = data_cliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
