https://www.udemy.com/course/blockchain-complete-hyperledger-fabric-development-bootcamp/learn/lecture/20925886\?start\=0\#overview



### Chaincode file sections (Low level)
- Import
- Struct
- Init function
- Invoke function
- Custom function
- Main function

### Chaincode file sections (High level)
- Import
- Struct
- Custom functions
- Main function


### Start network 
```
cd fabric-samples/test-network
./network.sh up createChannel -ca -s couchdb
```
### Install chain code
```
cd fabric-samples/test-network
./network.sh deployCC -ccn <NAME-OF-CHAINCODE> -ccp <PATH-TO-CHAINCODE-DIR> -ccl <CHAINCODE-LANGUAGE>
```