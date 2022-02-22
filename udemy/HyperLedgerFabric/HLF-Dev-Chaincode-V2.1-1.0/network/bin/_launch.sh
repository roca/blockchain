#PEER_MODE=dev
#Command=dev-init.sh -d 
#Generated: Tue Feb 22 15:33:02 UTC 2022 
docker-compose  -f ./compose/docker-compose.base.yaml    -f ./compose/docker-compose.dev.yaml      up -d --remove-orphans
