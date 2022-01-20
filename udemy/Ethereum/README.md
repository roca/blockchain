## Course link:
- https://www.udemy.com/course/blockchain-developer/

### blockchain: 
- https://ethereum-blockchain-developer.com
- https://ethereum-blockchain-developer.com/120-erc721-supply-chain-aisthisi/00-aisthisi-project-overview/?

### Online IDE:
- https://remix.ethereum.org


### TruffleSuite/Ganache
- https://trufflesuite.com/ganache/

Connect to http://127.0.0.1/7545 via Remix

### OpenZeppelin contract library: contracts/utils/math/SafeMath.sol
- https://github.com/OpenZeppelin/openzeppelin-contracts

path: /contracts/utils/math/SafeMath.sol


### GETH commands:
```
geth init ./genesis.json --datadir ./mychaindata
geth --datadir ./mychaindata --nodiscover
```
#### to attach:
```
geth attach ipc:/var/app/ex_private_network/mychaindata/geth.ipc
```

#### Coomon GETH commads:
```
	personal.newAccount();
	Passphrase: test123
	Repeat passphrase: test123
	"0x4efe4b21493dcda8794593712da04a75cf460b96"

	miner.setEtherbase(eth.accounts[0]);
	miner.start(1);

	web3.eth.getBalance(eth.accounts[0]);
```


#### Crowdsales:
- https://docs.openzeppelin.com/contracts/2.x/crowdsales