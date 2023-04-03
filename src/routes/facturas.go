package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/database/orm"
	"platcont/src/libraries/date"
	"platcont/src/libraries/library"
	"platcont/src/middleware"

	"github.com/gorilla/mux"
)

func RutasFacturas(r *mux.Router) {
	s := r.PathPrefix("/facturas").Subrouter()
	s.Handle("/list", middleware.Autentication(http.HandlerFunc(AllFactura))).Methods("GET")
	s.Handle("/list-factura", middleware.Autentication(http.HandlerFunc(FacturasDetalle))).Methods("GET")
	s.Handle("/update", middleware.Autentication(http.HandlerFunc(UpdateFactura))).Methods("PUT")

}

func AllFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	sessionID := r.Header.Get("Access-Token")

	id_fact := library.GetSession_key(sessionID, "id_user")

	data_facturaciones := orm.NewQuerys("facturas").Select().Where("id_clipd", "=", id_fact).Exec().All()
	if len(data_facturaciones) <= 0 {
		response.Msg = "Factura no encontrado"
		response.StatusCode = 300
		response.Status = "Factura no encontrado"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// controller.SessionMgr.SetSessionVal(controller.SessionID, "id_fact", data_facturaciones["id_fact"].(string))

	response.Data["facturas"] = data_facturaciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	sessionID := controller.SessionMgr.StartSession(w, r)

	id_fact := library.GetSession_key(sessionID, "id_user")
	if id_fact == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"id_clipd": id_fact}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Facturas_GetSchema()
	_Facturas := orm.SqlExec{}
	err = _Facturas.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Facturas.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Facturas.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func FacturasDetalle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	SessionId := r.Header.Get("Access-Token")
	id_clipd := library.GetSession_key(SessionId, "id_user")

	data_product_detail := orm.NewQuerys("productosdetalle").Select("months || '/' || years as periodo, s_impo, l_deta").Where("id_clipd", "=", id_clipd).Exec().All()

	var newFact []string

	for _, v := range data_product_detail {
		newFact = append(newFact, v["periodo"].(string))
	}
	// fmt.Println(newFact)

	data_client_product := orm.NewQuerys("clientproducts").Select("date_facture, s_impo").Where("id_clipd", "=", id_clipd).Exec().One()

	date_fact := date.GetDate(data_client_product["date_facture"].(string))
	date_now := date.GetDateLocation()

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	var data_facturaciones []map[string]interface{}
	var month = int64(12)

	impo := data_client_product["s_impo"]

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%02d", e)
			if library.IndexOf_String(newFact, year+"-"+month) == -1 {
				data_facturaciones = append(data_facturaciones, map[string]interface{}{
					"years":  year,
					"months": month,
					"s_impo": impo,
				})
			}
		}

		month_init = 1
	}
	response.Data["facturasDetail"] = data_facturaciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
