docker build -t hexbook-backend:latest .
docker service rm hexbook_backend
docker stack up -c ./service.yaml hexbook
docker service ls