package main

import (
	"akshikrm.github.io/pkg/templates"
	"context"
	"fmt"
	"github.com/a-h/templ"
	"net/http"
	"os"
)

func main() {

	if os.Args[1] == "build" {
		f, err := os.Create("docs/index.html")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = templates.Home().Render(context.Background(), f)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated docs/index.html")
	} else {
		fs := http.FileServer(http.Dir("./styles/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.Handle("/", templ.Handler(templates.Home()))

		fmt.Println("Listening on :3000")
		http.ListenAndServe(":3000", nil)
	}

}
