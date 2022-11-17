package main

import (
	// "io/ioutil"
	"net/http"
	"log"
	"time"
	"strconv"
	"encoding/json"
)


type Day struct {
	Name	string `json:"name"`
	Date	string `json:"date"`
	OffDay	bool `json:"isOffDay"`
}

type ResponseData struct {
	Papers	[]string `json:"papers"`
	Days	[]Day `json:"days"`
	Updated	time.Time
}

type API struct {
	Data map[int]*ResponseData
}


func NewAPI() *API {
	return &API{
		Data: make(map[int]*ResponseData),
	}
}

func (a *API) GetData(year int) *ResponseData {
	if data, ok := a.Data[year]; ok {
		// cache data in 7 days
		if time.Now().Sub(data.Updated).Hours() < 24 * 7 {
			return data
		}
	}

	a.Data[year] = a.FetchData(year)
	return a.Data[year]
}

func (a *API) FetchData(year int) *ResponseData {
	res, err := http.Get("https://raw.githubusercontent.com/NateScarlet/holiday-cn/master/"+ strconv.Itoa(year) +".json")
	if err != nil{
		log.Fatal(err)
	}

	defer res.Body.Close()

	var result ResponseData

	json.NewDecoder(res.Body).Decode(&result)
	result.Updated = time.Now()
	
	return &result
}


func (a *API) FillCalendar(c *Calendar, off bool) *Calendar {
	// get the next year
	year := time.Now().Year() + 1
	
	for y := year - 2; y <= year; y++ {
		data := a.GetData(y)

		for _, day := range data.Days {
			if day.OffDay != off {
				continue
			}
	
			t, _ := time.Parse("2006-01-02", day.Date)
			c.AddHoliday(day.Name, t, day.OffDay)
		}
	}

	return c
}

func (a *API) GetWorkCalendar() *Calendar {
	return a.FillCalendar(NewWorkCalendar(), false)
}

func (a *API) GetOffCalendar() *Calendar {
	return a.FillCalendar(NewOffCalendar(), true)
}
