package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func CheckBody(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorServer(w, err)
		return nil, err
	}

	data_request := make(map[string]interface{})
	json.Unmarshal(reqBody, &data_request)

	if len(data_request) <= 0 {
		ErrorsWaning(w, errors.New("No se obtuvo información "))
		return nil, errors.New("No se obtuvo información ")
	}

	return data_request, nil
}

func CheckBody_array(w http.ResponseWriter, r *http.Request) ([]map[string]interface{}, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorServer(w, err)
		return nil, err
	}

	var data_request [](map[string]interface{})
	json.Unmarshal(reqBody, &data_request)

	if len(data_request) <= 0 {
		ErrorsWaning(w, errors.New("No se obtuvo información "))
		return nil, errors.New("No se obtuvo información ")
	}

	return data_request, nil
}
