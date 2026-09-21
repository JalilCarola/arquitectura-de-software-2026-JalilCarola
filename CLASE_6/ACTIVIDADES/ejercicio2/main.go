package main

import (
	"fmt"
	"sync"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	workerCount  = 3
	processDelay = 260 * time.Millisecond
)

var productIDs = []int{200, 201, 202, 203, 204}

func updateProduct(workerID int, productID int) {
	fmt.Printf("Worker %d updated product %d\n", workerID, productID)
	time.Sleep(processDelay)
}

func processWithWorkers() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for workerID := 1; workerID <= workerCount; workerID++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for productID := range jobs {
				updateProduct(workerID, productID)
			}
		}(workerID)
	}

	for _, productID := range productIDs {
		jobs <- productID
	}
	close(jobs)

	wg.Wait()
}

func main() {
	startedAt := time.Now()
	processWithWorkers()
	fmt.Println("Time:", time.Since(startedAt))
}
