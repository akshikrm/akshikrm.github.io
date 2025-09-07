package main

import (
	"akshikrm.github.io/pkg/templates"
	"context"
	"fmt"
	"os"
)

func main() {
	// Example for a single page
	f, err := os.Create("public/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Replace 'templates.HomePage()' with your template function
	err = template.HomePage().Render(context.Background(), f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated public/index.html")
}
