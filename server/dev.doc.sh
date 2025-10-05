# generate swagger doc, should run at root to include all deps
cd /home/asun/project/fatcat/server
swag init -g cmd/main.go -o /home/asun/project/fatcat/server/docs --parseDependency --parseInternal