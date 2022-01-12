alert("hello world");

let web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));

// ### List accounts:
web3.eth.getAccounts().then(console.log);
let myContract = new web3.eth.Contract([
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
], "0xF3e8295d567c3d048D59058B378eb89AA97F8F72");

myContract.methods.myUint().call().then(result => console.log(result.toString()));

myContract.methods.setUint(123).send({from: web3.eth.accounts[0]}).then(console.log);

