package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		for {
			log.Println("App is working ok")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		time.Sleep(time.Minute)
		panic("fail")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	log.Println("Exiting on signal")
}
