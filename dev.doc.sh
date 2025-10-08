# generate swagger doc, should run at root to include all deps
swag init -g cmd/main.go -o ./docs --parseDependency --parseInternal