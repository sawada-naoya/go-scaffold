package main

import (
	"flag"
	"fmt"

	"go-scaffold/internal/generator"
)

func main() {
	name := flag.String("name", "", "Entity name (e.g. user)")
	flag.Parse()

	if *name == "" {
		fmt.Println("Please provide --name")
		return
	}

	layers := []string{"handler", "usecase", "service", "repository"}
	err := generator.Generate(*name, layers)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
