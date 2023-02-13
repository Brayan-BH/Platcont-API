package helper

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

// * Gestión de sesiones de sesión * /
type SessionMgr struct {
	mCookieName  string       // Nombre de cookie del cliente
	mLock        sync.RWMutex // exclusión mutua (para garantizar la seguridad del hilo)
	mMaxLifeTime int64        // Tiempo de recolección de basura

	mSessions map[string]*Session // puntero para guardar la sesión [sessionID] = session
}

// Cree un administrador de sesión (cookieName: el nombre de la cookie en el navegador; maxLifeTime: el ciclo de vida más largo)
func NewSessionMgr(cookieName string, maxLifeTime int64) *SessionMgr {
	mgr := &SessionMgr{mCookieName: cookieName, mMaxLifeTime: maxLifeTime, mSessions: make(map[string]*Session)}

	// Comienza a recuperar el tiempo
	go mgr.GC()

	return mgr
}

// En la página de inicio de la página de inicio, inicie la sesión
func (mgr *SessionMgr) StartSession(w http.ResponseWriter, r *http.Request) string {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	// Crea una nueva sesión independientemente de si estuvo originalmente presente o no
	newSessionID := url.QueryEscape(mgr.NewSessionID())

	// Guardar puntero
	var session *Session = &Session{mSessionID: newSessionID, mLastTimeAccessed: time.Now(), mValues: make(map[interface{}]interface{})}
	mgr.mSessions[newSessionID] = session
	// Deje que la cookie del navegador establezca el tiempo de vencimiento
	cookie := http.Cookie{Name: mgr.mCookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(mgr.mMaxLifeTime)}
	http.SetCookie(w, &cookie)

	return newSessionID
}

// Finalizar sesión
func (mgr *SessionMgr) EndSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.mCookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		mgr.mLock.Lock()
		defer mgr.mLock.Unlock()

		delete(mgr.mSessions, cookie.Value)

		// Deje que la cookie del navegador caduque inmediatamente
		expiration := time.Now()
		cookie := http.Cookie{Name: mgr.mCookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// Finalizar sesión
func (mgr *SessionMgr) EndSessionBy(sessionID string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	delete(mgr.mSessions, sessionID)
}

// Establecer el valor en la sesión
func (mgr *SessionMgr) SetSessionVal(sessionID string, key interface{}, value interface{}) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	if session, ok := mgr.mSessions[sessionID]; ok {
		session.mValues[key] = value
	}
}

// Obtenga el valor en la sesión
func (mgr *SessionMgr) GetSessionVal(sessionID string, key interface{}) (interface{}, bool) {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	if session, ok := mgr.mSessions[sessionID]; ok {
		if val, ok := session.mValues[key]; ok {
			return val, ok
		}
	}

	return nil, false
}

// Obtenga la lista de ID de sesión
func (mgr *SessionMgr) GetSessionIDList() []string {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	sessionIDList := make([]string, 0)

	for k, _ := range mgr.mSessions {
		sessionIDList = append(sessionIDList, k)
	}

	return sessionIDList[0:len(sessionIDList)]
}

// Juzgar la legitimidad de la cookie (debe juzgar la legitimidad cada vez que ingresa a una página)
func (mgr *SessionMgr) CheckCookieValid(w http.ResponseWriter, r *http.Request) string {
	var cookie, err = r.Cookie(mgr.mCookieName)

	if cookie == nil ||
		err != nil {
		return ""
	}

	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	sessionID := cookie.Value

	if session, ok := mgr.mSessions[sessionID]; ok {
		session.mLastTimeAccessed = time.Now() // Mientras juzga la validez, actualice la última hora de acceso
		return sessionID
	}

	return ""
}

// Actualizar el último tiempo de acceso
func (mgr *SessionMgr) GetLastAccessTime(sessionID string) time.Time {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	if session, ok := mgr.mSessions[sessionID]; ok {
		return session.mLastTimeAccessed
	}

	return time.Now()
}

// recuperación de GC
func (mgr *SessionMgr) GC() {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	for sessionID, session := range mgr.mSessions {
		// Eliminar la sesión que excede el límite de tiempo
		if session.mLastTimeAccessed.Unix()+mgr.mMaxLifeTime < time.Now().Unix() {
			delete(mgr.mSessions, sessionID)
		}
	}

	// Recuperación temporizada
	time.AfterFunc(time.Duration(mgr.mMaxLifeTime)*time.Second, func() { mgr.GC() })
}

// Crea una ID única
func (mgr *SessionMgr) NewSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano() // microsegundos
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}

//——————————————————————————
/*Sesión*/
type Session struct {
	mSessionID        string                      // ID único
	mLastTimeAccessed time.Time                   // Última hora de acceso
	mValues           map[interface{}]interface{} // Otros valores correspondientes (guarde algunos valores correspondientes al usuario, como los permisos de usuario)
}
