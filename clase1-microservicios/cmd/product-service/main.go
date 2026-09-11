package main

import (
	"log"
	"net/http"

	"clase1/starter/internal/products"
)

func main() {
	productRepo := products.NewRepository()
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /products", productHandler.List)
	mux.HandleFunc("GET /products/{id}", productHandler.GetByID)

	log.Println("product service listening on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
