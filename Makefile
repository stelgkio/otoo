run:
	@templ generate
	@go run main.go 

build:
	@go build -o bin/main main.go

templ:
	@templ generate -watch
build-css:
	./tailwindcss -i input.css -o assets/output.css --minify

dev-css:
	./tailwindcss -i input.css -o assets/output.css --watch=always
npx-css:
	npx tailwindcss -i input.css -o assets/output.css --watch