package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/libraries/library"
	"platcont/src/middleware"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasClientesProductos(r *mux.Router) {
	s := r.PathPrefix("/productos").Subrouter()
	s.Handle("/info", middleware.Autentication(http.HandlerFunc(ClientProducts))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(insertProduct))).Methods("POST")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(InsertProductsDetail))).Methods("POST")
	s.Handle("/update", middleware.Autentication(http.HandlerFunc(updateProduct))).Methods("PUT")

}

func insertProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	sessionID := controller.SessionMgr.StartSession(w, r)

	id_clie := library.GetSession_key(sessionID, "id_user")
	id_clipd := library.GetSession_key(sessionID, "id_user")

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_request["id_clie"] = id_clie
	data_request["id_clipd"] = id_clipd
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clientproducts_GetSchema()
	_Clientes_Products := go_basic_orm.SqlExec{}
	err = _Clientes_Products.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Clientes_Products.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		// controller.ErrorsWaning(w, errors.New("Error al Registrar producto"))
		return
	}

	returnData := _Clientes_Products.Data[0]
	response.Data = returnData
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	sessionID := controller.SessionMgr.StartSession(w, r)

	id_clie := library.GetSession_key(sessionID, "id_user")
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
	_Client_Products := go_basic_orm.SqlExec{}
	err = _Client_Products.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Client_Products.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Client_Products.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ClientProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	sessionID := r.Header.Get("Access-Token")

	id_clipd := library.GetSession_key(sessionID, "id_user")

	data_client_product, _ := new(go_basic_orm.Querys).NewQuerys("clientproducts").Select("data_base,date_facture,host,id_clie,id_clipd,modulos, multi,users").Where("id_clie", "=", id_clipd).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(data_client_product) <= 0 {
		response.Msg = "Producto no encontrado"
		response.StatusCode = 300
		response.Status = "Producto no encontrado"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// controller.SessionMgr.SetSessionVal(controller.SessionID, "id_clipd", data_client_product["id_clipd"].(string))

	response.Data["productClient"] = data_client_product
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
