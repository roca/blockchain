
// ### From Node repl:
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider("http://ganache:8545"));

// ### List accounts:
web3.eth.getAccounts().then(console.log);

// ### Get balance:
web3.eth.getBalance("0x263Eb2FC832a480721501320315fFeE3A68Fe35e").then(console.log);

web3.eth.getBalance("0x668E8DF2fEA7039A3D2Da978E9B721c9C8394C58").then(function(result) {
	  console.log(web3.utils.fromWei(result, "ether"));
});

// ### Transfer:
web3.eth.sendTransaction({
	from: "0x263Eb2FC832a480721501320315fFeE3A68Fe35e", 
	to:"0x668E8DF2fEA7039A3D2Da978E9B721c9C8394C58", 
	value: web3.utils.toWei("1", "ether")}).then(console.log);
