package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"platcont/src/controller"
	"platcont/src/library/sqlquery"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type jwtclaim struct {
	User  string `json:"user"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func RutasAuth(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")

}

func auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-Json")
	response := controller.NewResponseManager()

	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	body := make(map[string]interface{})
	json.Unmarshal(req_body, &body)

	dataUser := sqlquery.NewQuerys("users").Select().Where("email", "=", body["email"])
	if len(dataUser) <= 0 {
		response.Msg = "Usuario y Contraseña Inconrrecto"
		response.StatusCode = 300
		response.Status = "Usuario y Contraseña Inconrrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["password"].(string)), []byte(body["password"].(string)))
	if err != nil {
		response.Msg = "Usuario y Contraseña Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	//Tokens
	var key_token interface{}
	key_token = []byte("supervisor")
	claims := jwtclaim{
		dataUser["users"].(string),
		dataUser["email"].(string),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * (60 * 24)).Unix(),
			Issuer:    "pdt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err_token := token.SignedString(key_token)
	if err_token != nil {
		response.Msg = "Error signing" + err_token.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["token"] = token_string
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
