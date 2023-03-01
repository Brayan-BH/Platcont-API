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

	"github.com/gorilla/mux"
)

func RutasFacturas(r *mux.Router) {
	s := r.PathPrefix("/facturas").Subrouter()
	s.Handle("/list", (http.HandlerFunc(allFacturas))).Methods("GET")
	s.Handle("/info/{id_clie}", (http.HandlerFunc(oneProduct))).Methods("GET")
	s.Handle("/list-factura/{id_clipd}", (http.HandlerFunc(facturasDetalle))).Methods("GET")
	s.Handle("/create/{id_clipd}", (http.HandlerFunc(regDetalleFactura))).Methods("POST")

}

func allFacturas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_factura := orm.NewQuerys("Facturas").Select().Exec().All()

	if len(data_factura) <= 0 {
		controller.ErrorsSuccess(w, errors.New("No se encontraron resultados para la consulta"))
		return
	}

	response.Data["facturas"] = data_factura
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func regDetalleFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	// params := mux.Vars(r)
	// id_clipd := params["id_clipd"]
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Facturas_GetSchema()
	_Facturas := orm.SqlExec{}
	err = _Facturas.New(data_insert, table).Insert(schema)
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

//TODO: PRODUCTOS DETALLE

func facturasDetalle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_pddt := params["id_pddt"]

	if id_pddt == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	pddt_detail := orm.NewQuerys("productosdetalle").Select("years || '-' || months as periodo,l_deta,id_clipd,id_pddt,s_impo").Where("id_pddt", "=", id_pddt).Exec().All()

	var newFact []string
	pago := map[string]interface{}{
		"if":      false,
		"periodo": "",
		"s_impo":  0,
	}

	for _, v := range pddt_detail {
		newFact = append(newFact, v["periodo"].(string))

		// if v["k_stad"].(int64) == 1 {

		// 	pagoCuenta["if"] = true
		// 	pagoCuenta["periodo"] = v["periodo"]
		// 	pagoCuenta["s_impo"] = v["s_impo"]
		// }
		if v["s_impo"] == true {
			pago["s_impo"] = v["s_impo"]
		}
	}
	// fmt.Println(newFact)

	id_clipd := params["id_clipd"]

	if id_clipd == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_fact_detail := orm.NewQuerys("Facturas_Detalle").Select("n_item", "c_prod", "s_impo", "s_desc", "s_igv", "s_tota", "l_peri").Where("id_clipd", "=", id_clipd).Exec().One()
	date_fact := date.GetDate(data_fact_detail["periodo"].(string))
	date_now := date.GetDateLocation()

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	var data_facturaciones []map[string]interface{}
	var month = int64(12)

	impo := data_fact_detail["s_impo"].(float64)

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%d", e)
			if library.IndexOf_String(newFact, year+"-"+month) == -1 {
				data_facturaciones = append(data_facturaciones, map[string]interface{}{
					"year":      year,
					"month":     month,
					"s_impo":    impo,
					"last_amor": 0,
				})
			} else {
				if pago["if"].(bool) == true {
					data_facturaciones = append(data_facturaciones, map[string]interface{}{
						"year":      year,
						"month":     month,
						"s_impo":    impo,
					})
				}
			}
		}

		month_init = 1
	}

	response.Data["detalle_fact"] = data_facturaciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
