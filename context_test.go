package golangcontext

import (
	"context"
	"fmt"
	"testing"
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
