package main

import (
	"context"
	"fmt"
	"os"

	"akshikrm.github.io/pkg/templates"
)

func main() {
	// Example for a single page
	f, err := os.Create("docs/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Replace 'templates.HomePage()' with your template function
	err = templates.Home().Render(context.Background(), f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated docs/index.html")
}
