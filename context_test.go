package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestCreateContextWithValue(t *testing.T) {

	// Creating root context (Super parent)
	contextA := context.Background()

	// Creating childs
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	// Print out the context
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Get data from context
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))

	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("c"))

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		// it's makes a goroutine leaks
		for {

			// Menambahkan sebuah pengecekan terhadap context
			select {

			// if ctx.Done() return a value, looping will stop and gorouting will stop also
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	// create parent context
	parent := context.Background()

	// create child context with cancel
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	// execute the cencel(), so the done() function will return a value, and goroutine will stop
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}
