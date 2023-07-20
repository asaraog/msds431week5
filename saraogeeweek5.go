package main

import (
	"fmt"
	"os"
)

func main() {
	os.Create("items.jl")
	fmt.Println()
}
