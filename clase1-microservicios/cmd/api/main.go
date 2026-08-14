package main

import (
	"log"
	"net/http"

	"clase1/starter/internal/orders"
	"clase1/starter/internal/products"
)

func main() {
	productRepo := products.NewRepository()
	orderRepo := orders.NewRepository()

	productService := products.NewService(productRepo)
	orderService := orders.NewService(orderRepo, productRepo)

	productHandler := products.NewHandler(productService)
	orderHandler := orders.NewHandler(orderService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /products", productHandler.List)
	mux.HandleFunc("GET /products/{id}", productHandler.GetByID)
	mux.HandleFunc("GET /orders", orderHandler.List)
	mux.HandleFunc("POST /orders", orderHandler.Create)

	log.Println("starter api listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
