package main

import (
	"fmt"

	"github.com/manoj-gupta/go1_18beta1/fuzz"
	"github.com/manoj-gupta/go1_18beta1/generics"
)

func main() {
	// generics
	fmt.Println("Testing Generics .....")
	generics.Run()

	// fuzz
	fmt.Println("Testing Fuzzing ......")
	fuzz.Run()
}
