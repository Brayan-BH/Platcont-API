package routes

import (
	"encoding/json"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/libraries/library"
	"platcont/src/middleware"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutasProductosDetalle(r *mux.Router) {
	s := r.PathPrefix("/productos-detail").Subrouter()
	s.Handle("/info", middleware.Autentication(http.HandlerFunc(ProductsDetail))).Methods("GET")
	s.Handle("/create", middleware.Autentication(http.HandlerFunc(InsertProductsDetail))).Methods("POST")

}

func InsertProductsDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	sessionID := r.Header.Get("Access-Token")

	id_clipd := library.GetSession_key(sessionID, "id_user")
	id_fact := library.GetSession_key(sessionID, "id_user")

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_request["id_clipd"] = id_clipd
	data_request["id_fact"] = id_fact
	data_insert = append(data_insert, data_request)

	schema, table := tables.Productosdetalle_GetSchema()
	_Products_Detail := go_basic_orm.SqlExec{}
	err = _Products_Detail.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Products_Detail.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	returnData := _Products_Detail.Data[0]
	response.Data = returnData
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ProductsDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	sessionID := r.Header.Get("Access-Token")

	id_pddt := library.GetSession_key(sessionID, "id_user")

	data_product_detail, _ := new(go_basic_orm.Querys).NewQuerys("productosdetalle").Select().Where("id_clipd", "=", id_pddt).Exec(go_basic_orm.Config_Query{Cloud: true}).All()

	if len(data_product_detail) <= 0 {
		response.Msg = "Detalle de producto no encontrado"
		response.StatusCode = 300
		response.Status = "Detalle de producto no encontrado"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["productDetail"] = data_product_detail
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
