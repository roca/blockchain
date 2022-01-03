// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract FunctionsExample {
	mapping(address => uint64) public balanceReceived;

	address payable owner;

	constructor() {
		owner = payable(msg.sender);
	}

	function getOwner() public view returns(address) {
		return owner;
	}

	function convertWeiToEther(uint _amountInWei) public pure returns(uint) {
		return _amountInWei / 1 ether;
	}

	function destroySmartContract() public {
		require(payable(msg.sender) == owner, "You are not the owner");
		selfdestruct(owner);
	}

	function receiveMoney() public payable {
		assert(balanceReceived[msg.sender] + uint64(msg.value) >= balanceReceived[msg.sender]);
		balanceReceived[msg.sender] += uint64(msg.value);
	}

	function withdrawMoney(address payable _to, uint64 _amount) public {
		require( _amount <= balanceReceived[msg.sender], "You don't have enough ether");
		assert(balanceReceived[msg.sender] >= balanceReceived[msg.sender] - _amount);
		balanceReceived[msg.sender] -= _amount;
		_to.transfer(_amount);
		
	}

	// fallback() external payable {
	// 	require(msg.value == 0, "You can't send ether to this contract");
	// 	receiveMoney();
	// } 

	receive() external payable {
        require(msg.value > 0, "You can't send 0 ether to this contract");
		receiveMoney();
    	}
}