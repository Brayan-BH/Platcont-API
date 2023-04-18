package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"platcont/src/controller"
	"platcont/src/database/models/tables"
	"platcont/src/libraries/library"
	"platcont/src/middleware"

	"github.com/deybin/go_basic_orm"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func RutasAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
	s.HandleFunc("/logout", logout).Methods("PUT")
	s.HandleFunc("/first-step", registerFirst).Methods("POST")
	s.Handle("/second-step", middleware.Autentication(http.HandlerFunc(registerSecond))).Methods("POST")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func logout(w http.ResponseWriter, r *http.Request) {
	// Finalizar la sesión del usuario
	cookies, _ := r.Cookie("cookie-token")
	sessionID := cookies.Value
	controller.SessionMgr.EndSessionBy(sessionID)
	err := controller.NewResponseManager()

	if err != nil {
		// Si hay un error al finalizar la sesión, mostrar un mensaje de error
		// Si todo va bien, mostrar una respuesta exitosa
		response := controller.NewResponseManager()
		response.Msg = "Sesión cerrada exitosamente"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	// Get the request body
	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//objeto map
	body := make(map[string]interface{})

	json.Unmarshal(req_body, &body)

	dataUser, _ := new(go_basic_orm.Querys).NewQuerys("users").Select().Where("email", "=", body["email"]).Exec(go_basic_orm.Config_Query{Cloud: true}).One()

	if len(dataUser) <= 0 {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contraseña Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["password_admin"].(string)), []byte(body["password_admin"].(string)))
	if err != nil {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	SessionID := controller.SessionMgr.StartSession(w, r)

	// Establecer el valor "login" en la sesión
	controller.SessionMgr.SetSessionVal(SessionID, "login", true)
	controller.SessionMgr.SetSessionVal(SessionID, "id_user", dataUser["id_user"].(string))

	returnData := dataUser
	delete(returnData, "id_user")
	delete(returnData, "password_admin")
	delete(returnData, "password")
	response.Data["users"] = returnData
	response.Data["cookie_token"] = SessionID
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func registerFirst(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Users_GetSchema()
	clientes := go_basic_orm.SqlExec{}
	err = clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = clientes.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	returnData := clientes.Data[0]
	delete(returnData, "id_user")
	delete(returnData, "password_admin")
	delete(returnData, "password")
	response.Data = returnData

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func registerSecond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	SessionID := controller.SessionMgr.StartSession(w, r)

	id_clie := library.GetSession_key(SessionID, "id_user")

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_request["id_clie"] = id_clie
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clients_GetSchema()
	clientes := go_basic_orm.SqlExec{}
	err = clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = clientes.Exec("Platcont")
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	returnData := clientes.Data[0]
	response.Data = returnData
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
