package routes

import (
	"encoding/json"
	"net/http"
	"platcont/src/controller"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
)

func RutaVersiones(r *mux.Router) {
	s := r.PathPrefix("/version").Subrouter()
	s.Handle("/last", (http.HandlerFunc(LastVersion))).Methods("GET")

}

func LastVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	// Obtener la versión más reciente ordenando por fecha descendente
	data_versiones, _ := new(go_basic_orm.Querys).NewQuerys("versiones").
		Select("c_vers, id_file, l_deta").
		OrderBy("f_digi DESC").
		Limit(1).
		Exec(go_basic_orm.Config_Query{Cloud: true}).
		One()

	if len(data_versiones) <= 0 {
		response.Msg = "Version no encontrada"
		response.StatusCode = 300
		response.Status = "Version no encontrada"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["last_version"] = data_versiones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
