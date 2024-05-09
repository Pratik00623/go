// Create a program that performs a lengthy calculation and can be canceled on user demand. Use the 
// context package to handle cancellation gracefully, demonstrating the ability to stop ongoing 
// operations.



package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func calculateFactorial(ctx context.Context, n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Calculation canceled by user.")
			return nil
		default:
			result.Mul(result, big.NewInt(int64(i)))
			// Introduce a small delay to simulate a lengthy calculation
			time.Sleep(500 * time.Millisecond)
		}
	}
	return result
}

func main() {
	var n int
	fmt.Print("Enter a positive integer: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil || n < 0 {
		fmt.Println("Invalid input. Please enter a valid positive integer.")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle cancellation via Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nReceived interrupt signal. Canceling calculation...")
		cancel()
	}()

	factorialResult := calculateFactorial(ctx, n)
	if factorialResult != nil {
		fmt.Printf("The factorial of %d is %s\n", n, factorialResult.String())
	}
}
