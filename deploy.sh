docker build -t hub-hexbook:latest .
docker service rm hub_hexbook
docker stack up -c ./service.yaml hub
echo [LOG] waiting 5 seconds to deploy
sleep 5
echo [LOG] done, listing up services
docker service ls