run:
	@templ generate
	@go run cmd/main.go 

build:
	@go build -o bin/main cmd/main.go
build-css:
	./tailwindcss -i input.css -o assets/output.css --minify

dev-css:
	./tailwindcss -i input.css -o assets/output.css --watch=always
npx-css:
	npx tailwindcss -i input.css -o assets/output.css --watch