package main

import (
	"log"
	"github.com/pionus/arry"
	"github.com/pionus/arry/middlewares"
)

func main() {
	a := arry.New()
	a.Use(middlewares.Gzip)
	a.Use(middlewares.Logger())
	a.Use(middlewares.Panic)

	api := NewAPI()

	router := a.Router()

	router.Get("/work", func(ctx arry.Context) {
		work := api.GetWorkCalendar()
		ctx.Response().Code = 200
		work.Render(ctx.Response().Writer)
	})

	router.Get("/off", func(ctx arry.Context) {
		off := api.GetOffCalendar()
		ctx.Response().Code = 200
		off.Render(ctx.Response().Writer)
	})

	log.Printf("Listening at :80")
	err := a.Start(":80")

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

	log.Printf("shutdown")
}