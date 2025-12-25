INPUT_CSS = ./input.css
OUTPUT_CSS = ./web/static/css/output.css
SERVER_PORT = 3000

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