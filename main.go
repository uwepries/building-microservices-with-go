package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/uwepries/building-microservices-with-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(l)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", productsHandler)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
