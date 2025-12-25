INPUT_CSS = ./input.css
OUTPUT_CSS = ./web/static/css/output.css
SERVER_PORT = 3000

install:
	@echo "Installing go..."
	@sudo apt update && sudo apt install golang-go

	@echo "Installing tailwind..."
	@echo "Setting daisyUI..."
	@curl -sL daisyui.com/fast | bash

	@echo "Installing templ..."
	@go get -tool github.com/a-h/templ/cmd/templ@latest

	@echo "Installing gsap..."
	@curl -O https://raw.githubusercontent.com/greensock/GSAP/refs/heads/master/dist/gsap.min.js
	@curl -O https://raw.githubusercontent.com/greensock/GSAP/refs/heads/master/dist/ScrollTrigger.min.js
	@mv gsap.min.js ScrollTrigger.min.js web/static/js/gsap

	@echo "Installing wgo..."
	@go install github.com/bokwoon95/wgo@latest

	@echo "Installing make build tool..."
	@sudo apt update && sudo apt install build-essential

# Generate css from from html class 
tailwind:
	@echo "Watching Tailwind..."
	@./tailwindcss -i $(INPUT_CSS) -o $(OUTPUT_CSS) --watch

# Templ
templ-watch:
	@echo "Generating html from templ files..."
	@go tool templ generate --watch

# Run server
run:
	@go run .

# Development - runs all watchers in parallel
dev:
	@echo "Starting development environment..."
	@bash -c '\
		trap "kill 0" EXIT; \
		go tool templ generate --watch --proxy="http://localhost:$(SERVER_PORT)" --cmd="go run ." & \
		./tailwindcss -i $(INPUT_CSS) -o $(OUTPUT_CSS) --watch & \
		wait \
	'

# Alternative: Run separate commands
dev-css:
	@echo "Starting Tailwind watcher..."
	@./tailwindcss -i $(INPUT_CSS) -o $(OUTPUT_CSS) --watch

dev-server:
	@echo "Starting wgo server..."
	@wgo -file=.go -file=.templ -xfile=_templ.go go run .

.PHONY: tailwind templ-gen run dev dev-css dev-server