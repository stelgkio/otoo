run:
	@templ generate
	@go run main.go 

build:
	@go build -o bin/main main.go

templ:
	@templ generate -watch
build-css:


dev-css:

npx-css:
