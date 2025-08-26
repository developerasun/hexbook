docker build -t fatcat-frontend:latest .
docker service rm fatcat_frontend fatcat_clienttunnel
docker stack up -c ./service.yaml fatcat
docker service ls