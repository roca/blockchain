#PEER_MODE=net
#Command=dev-init.sh -e 
#Generated: Mon Feb 28 18:18:38 UTC 2022 
docker-compose  -f ./compose/docker-compose.base.yaml      -f ./compose/docker-compose.explorer.yaml    up -d --remove-orphans
