#PEER_MODE=net
#Command=dev-init.sh -s 
#Generated: Fri Mar 18 13:45:36 UTC 2022 
docker-compose  -f ./compose/docker-compose.base.yaml     -f ./compose/docker-compose.couchdb.yaml     up -d --remove-orphans
