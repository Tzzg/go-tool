package main

import (
	"context"
	"github.com/Tzzg/go-tool/Data"
	"github.com/Tzzg/go-tool/worker_pool"
	"log"
	"time"
)

func main() {
	// test Data
	a := Data.StrToMd5("asdsa", 10)
	log.Printf(a)

	// test work pool
	wp := worker_pool.NewWorkerPool(context.TODO(), 16, 10000)

	add := wp.TryAdd(func(ctx context.Context) {
		time.Sleep(3)
		log.Println("echo 1")
	})
	if add == false {
		log.Printf("wp.TryAdd false")
	}
	log.Println("WaitAndClose")
	wp.WaitAndClose()
}
