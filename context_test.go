package main

import (
	"context"
	"fmt"
	"testing"
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
