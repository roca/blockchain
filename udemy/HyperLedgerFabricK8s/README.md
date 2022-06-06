https://www.udemy.com/course/hyperledger-fabric-on-kubernetes-complete-guide


### Mounting filesystems

- https://minikube.sigs.k8s.io/docs/commands/start/

#### Stack overflow:
- https://stackoverflow.com/questions/48534980/mount-local-directory-into-pod-in-minikube
- https://minikube.sigs.k8s.io/docs/handbook/persistent_volumes/


```
minikube mount $HOME/GitHub/Repos/blockchain/udemy/HyperLedgerFabricK8s/data:/data
minikube start --driver=hyperkit --mount-string="$HOME/GitHub/Repos/blockchain/udemy/HyperLedgerFabricK8s/data:/data" --mount --nodes 4 -p multinode-demo
minikube start --driver=virtualbox  --mount-string="$HOME/GitHub/Repos/blockchain/udemy/HyperLedgerFabricK8s/data:/data" --mount --nodes 4 -p multinode-demo
```