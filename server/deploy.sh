docker service rm fatcat_backend
docker stack up -c ./service.yaml fatcat
docker service ls
