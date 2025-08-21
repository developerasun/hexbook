# docker build -t fatcat-backend:latest .
docker service rm fatcat_backend fatcat_tunnel
docker stack up -c ./service.yaml fatcat
docker service ls