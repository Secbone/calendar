package main

import (
	"io/ioutil"
	"net/http"
	"log"
)


type Holiday struct {
	Holiday	bool `json:"holiday"`
	Name	string `json:"name"`
	Wage	int `json:"wage"`
	Date	string `json:"date`
}

type ResponseData struct {
	Code	int `json:"code"`
	Holiday	map[string]Holiday `json:"holiday"`
}


func GetHolidays() {
	res, err := http.Get("http://timor.tech/api/holiday/year/")
	if err != nil{
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Print(string(body))
}