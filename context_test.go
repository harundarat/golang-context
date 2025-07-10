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
