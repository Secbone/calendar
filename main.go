package main

import (
	"os"
	"os/signal"
	"fmt"
	"time"
	"context"
)

func main() {
	ctx := context.Background()
	worker := NewScheduler()

	worker.Add(ctx, testJob, time.Second * 2)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<- quit
	worker.Stop()
}

func testJob(ctx context.Context) {
	fmt.Println("job called")
	GetHolidays()
}