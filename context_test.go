package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("Background context:", background)

	todo := context.TODO()
	fmt.Println("TODO context:", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f")) // Output: F
	fmt.Println(contextF.Value("c")) // Output: C
	fmt.Println(contextF.Value("b")) // Output: <nil> -> contextB is not an ancestor of contextF

	fmt.Println(contextA.Value("b")) // Output: <nil> -> contextB is not an ancestor of contextA

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulate some work
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutines before:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)
	fmt.Println("Total goroutines after:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	cancel()                    // Send cancel signal to context to stop the goroutine
	time.Sleep(2 * time.Second) // Wait for goroutine to finish
	fmt.Println("Total goroutines after cancel:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutines before:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel() // Ensure the context is cancelled after use
	destination := CreateCounter(ctx)
	fmt.Println("Total goroutines after:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	time.Sleep(2 * time.Second) // Wait for goroutine to finish
	fmt.Println("Total goroutines after timeout:", runtime.NumGoroutine())
}
