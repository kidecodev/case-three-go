package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

//GetDataFromRajaOngkir ...
func GetDataFromRajaOngkir(api string) []byte {
	res, err := http.Get(api)
	if err != nil {
		log.Fatalln("Error get data")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error get data")
	}

	return data;
}
