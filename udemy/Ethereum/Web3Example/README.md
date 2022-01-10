```
npm install --save web3
```


### From Node repl:
```
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider("http://ganache:8545"))
```

### List accounts:
```
web3.eth.getAccounts().then(console.log);
```
### Get balance:
```
web3.eth.getBalance("0x060C327D30D46B5b252b11A46A9c6f25e80087b8").then(console.log);
```
