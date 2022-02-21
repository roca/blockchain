#PEER_MODE=net
#Command=dev-init.sh -e 
#Generated: Mon Feb 21 10:01:23 UTC 2022 
docker-compose  -f ./compose/docker-compose.base.yaml      -f ./compose/docker-compose.explorer.yaml    up -d --remove-orphans
