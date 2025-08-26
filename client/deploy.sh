docker build -t fatcat-frontend:latest .
docker service rm fatcat_frontend
docker stack up -c ./service.yaml fatcat
docker service ls