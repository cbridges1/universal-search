package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"universal-search/internal/models"
)

var myClient = &http.Client{Timeout: 20 * time.Second}

func Get(url string) []byte {

	response, er := myClient.Get(url)

	if er != nil {
		log.Fatal(er)
		// return er.Error()
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		// return errError(
	}

	return responseData
}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func ReturnJsonResponseJ(res http.ResponseWriter, httpCode int, resMessage models.Response) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	json.NewEncoder(res).Encode(resMessage)
}
