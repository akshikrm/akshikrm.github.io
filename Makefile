all: generate dev

generate:
	templ generate
dev: 
	go run main.go 

build: generate
	rm ./docs/main.css
	cp ./styles/main.css ./docs/main.css
	go run main.go build


	
