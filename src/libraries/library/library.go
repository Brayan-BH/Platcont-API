package library

import (
	"fmt"
	"platcont/src/controller"
	"reflect"
	"strings"
)

//typeof
//recibe un valor interface que no se reconoce su tipo y devuelve un string
func GetSession_key(SessionID string, key string) interface{} {
	data, err := controller.SessionMgr.GetSessionVal(SessionID, key)
	if !err {
		fmt.Println("Error session ID:"+SessionID+" key:("+key+"):", err)
		fmt.Println("lista:", controller.SessionMgr.GetSessionIDList())
		return ""
	}
	return data
}
func InterfaceToString(params ...interface{}) string {
	typeValue := reflect.TypeOf(params[0]).String()
	value := params[0]
	valueReturn := ""
	if strings.Contains(typeValue, "string") {
		toSql := false
		if len(params) == 2 && reflect.TypeOf(params[1]).Kind() == reflect.Bool {
			toSql = params[1].(bool)
		}

		if toSql {
			valueReturn = fmt.Sprintf("'%s'", value)
		} else {
			valueReturn = fmt.Sprintf("%s", value)
		}
	} else if strings.Contains(typeValue, "int") {
		valueReturn = fmt.Sprintf("%d", value)
	} else if strings.Contains(typeValue, "float") {
		valueReturn = fmt.Sprintf("%f", value)
	} else if strings.Contains(typeValue, "bool") {
		valueReturn = fmt.Sprintf("%t", value)
	}
	return valueReturn
}

func IndexOf_String(arreglo []string, search string) int {
	for indice, valor := range arreglo {
		if valor == search {
			return indice
		}
	}
	// -1 porque no existe
	return -1
}

func IndexOf_String_Map(arreglo []map[string]interface{}, key, search string) int {
	for indice, valor := range arreglo {
		if valor[key] == search {
			return indice
		}
	}
	// -1 porque no existe
	return -1
}
