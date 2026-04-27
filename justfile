port := "8080"
log_level := "INFO"
environment := "local"

default: run

dev:
    air

build:
    go build -o build/main .

run:
    PORT={{port}} LOG_LEVEL={{log_level}} ENVIRONMENT={{environment}} go run .
