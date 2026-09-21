package main

import (
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	firstStoreName  = "Tienda 10 A"
	firstStorePrice = 1730
	firstStoreDelay = 820 * time.Millisecond

	secondStoreName  = "Tienda 10 B"
	secondStorePrice = 1510
	secondStoreDelay = 380 * time.Millisecond
)

type Price struct {
	Store string
	Value int
}

func searchFirstStore() Price {
	time.Sleep(firstStoreDelay)
	return Price{Store: firstStoreName, Value: firstStorePrice}
}

func searchSecondStore() Price {
	time.Sleep(secondStoreDelay)
	return Price{Store: secondStoreName, Value: secondStorePrice}
}

func comparePrices() []Price {
	results := make(chan Price, 2)

	go func() {
		results <- searchFirstStore()
	}()
	go func() {
		results <- searchSecondStore()
	}()

	first := <-results
	second := <-results

	return []Price{first, second}
}

func main() {
	startedAt := time.Now()
	for _, price := range comparePrices() {
		fmt.Printf("%s: $%d\n", price.Store, price.Value)
	}
	fmt.Println("Time:", time.Since(startedAt))
}
