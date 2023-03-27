package middleware

import (
	"net/http"
	"platcont/src/controller"
	"platcont/src/libraries/library"

	"github.com/gorilla/mux"
)

func Autentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()

		sessionID := r.Header.Get("Access-Token")

		// sessionID := controller.SessionMgr.CheckCookieValid(w, r)
		if sessionID == "" {
			response.Msg = "Debe Iniciar Session"
			response.Status = "Error"
			response.StatusCode = 401 //inautorizado
			w.Header().Set("Content-Type", "Aplication-Json")
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		session := library.GetSession_key(sessionID, "login")
		if session != nil {
			if session == true {
				next.ServeHTTP(w, r)
			} else {
				response.Msg = "Debe Iniciar Session"
				response.Status = "Error"
				response.StatusCode = 401 //inautorizado
				w.Header().Set("Content-Type", "Aplication-Json")
				w.WriteHeader(http.StatusAccepted)
			}
		} else {
			response.Msg = "Debe Iniciar Session"
			response.Status = "Error"
			response.StatusCode = 401 //inautorizado
			w.Header().Set("Content-Type", "Aplication-Json")
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")
			next.ServeHTTP(w, req)
		})
}

func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(MiddlewareCors)
}
