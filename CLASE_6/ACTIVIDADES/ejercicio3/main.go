package main

import (
	"context"
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	productName    = "Producto 10"
	reviewsScore   = 40
	reviewsDelay   = 1700 * time.Millisecond
	timeoutLimit   = 350 * time.Millisecond
	defaultReviews = 0
)

func getReviews() int {
	time.Sleep(reviewsDelay)
	return reviewsScore
}

func getProductReviews() int {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutLimit)
	defer cancel()

	result := make(chan int, 1)
	go func() {
		result <- getReviews()
	}()

	select {
	case reviews := <-result:
		return reviews
	case <-ctx.Done():
		return defaultReviews
	}
}

func main() {
	startedAt := time.Now()
	reviews := getProductReviews()
	fmt.Printf("%s reviews: %d\n", productName, reviews)
	fmt.Println("Time:", time.Since(startedAt))
}
