package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longRunningTask(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d started\n", id)

	select {
	case <-time.After(time.Duration(1+rand.Intn(5)) * time.Second):
		fmt.Printf("Task %d finished\n", id)
	case <-ctx.Done():
		fmt.Printf("Task %d canceled\n", id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	// Создаем контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Запускаем 5 горутин
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go longRunningTask(ctx, i, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Program finished")
}
