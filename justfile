port := "8080"
log_level := "INFO"
environment := "local"

default: run

css:
    tailwindcss -i assets/styles/input.css -o static/styles/tailwind.css --minify

css-watch:
    tailwindcss -i assets/styles/input.css -o static/styles/tailwind.css --watch

dev:
    air

build: css
    go build -o build/main .

run: css
    PORT={{port}} LOG_LEVEL={{log_level}} ENVIRONMENT={{environment}} go run .
