INPUT_CSS = input.css
OUTPUT_CSS = web/static/css/output.css

# generate css file from tailwind
tailwind:
	@echo "Building css files..." 
	@./tailwindcss -i $(INPUT_CSS) -o $(OUTPUT_CSS) --watch 

# generate go file from temple
templ:
	@echo "Generating go code from templ files..." 
	@go tool templ generate -watch 

go_run:
	@clear
	@echo "Running server..."
	@go run .

.PHONY: tailwind templ

# watch changes and generate files
watch: tailwind templ

# start server
start: go_run