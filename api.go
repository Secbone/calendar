package main

import (
	// "io/ioutil"
	"net/http"
	"log"
	"encoding/json"
)


type Holiday struct {
	Name	string `json:"name"`
	Date	string `json:"date`
	OffDay	bool `json:"isOffDay"`
}

type ResponseData struct {
	Papers	[]string `json:"papers"`
	Days	[]Holiday `json:"days"`
}


func GetHolidays() []Holiday {
	res, err := http.Get("https://raw.githubusercontent.com/NateScarlet/holiday-cn/master/2020.json")
	if err != nil{
		log.Fatal(err)
	}

	defer res.Body.Close()

	var result ResponseData

	json.NewDecoder(res.Body).Decode(&result)

	for _, h := range result.Days {
		log.Println(h.Name)
	}

	return result.Days
}