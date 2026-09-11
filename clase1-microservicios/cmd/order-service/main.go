package main

import (
	"log"
	"net/http"

	"clase1/starter/internal/orders"
)

func main() {
	productClient := orders.NewHTTPProductClient("http://localhost:8081")
	orderRepo := orders.NewRepository()
	orderService := orders.NewService(orderRepo, productClient)
	orderHandler := orders.NewHandler(orderService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /orders", orderHandler.List)
	mux.HandleFunc("POST /orders", orderHandler.Create)

	log.Println("order service listening on :8082")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
