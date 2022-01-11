// ### From Node repl:
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider("http://ganache:8545"));

// ### List accounts:
web3.eth.getAccounts().then(console.log);

// ### Call:
let funcSignitureHash = web3.utils.sha3("myUint()").substr(0, 10);

web3.eth.call({
	from:"0x263Eb2FC832a480721501320315fFeE3A68Fe35e", 
	to:"0x1134b2195abe82CCab111C26F8b6ac7E2AEd0420", 
	data: funcSignitureHash}).then(console.log);

console.log("Hash of function signiture: " + funcSignitureHash);

let contract = new web3.eth.Contract([
	{
		"inputs": [],
		"name": "myUint",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_myUint",
				"type": "uint256"
			}
		],
		"name": "setUint",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
], "0x1134b2195abe82CCab111C26F8b6ac7E2AEd0420");

contract.methods.myUint().call().then(console.log);
contract.methods.setUint(59).send({from:"0x263Eb2FC832a480721501320315fFeE3A68Fe35e"}).then(console.log);
// 0x06540f7e