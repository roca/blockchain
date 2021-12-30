// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract StartStopUpdateExample {
	address owner;

	bool public paused;

	constructor() {
		owner = msg.sender;
	}

	function sendMoney() public payable {
	}

	function setPaused(bool _paused) public {
		require(msg.sender == owner, "Only owner can pause the contract");
		paused = _paused;
	}

	function withdrawAllMoney(address payable _to) public {
		require(msg.sender == owner, "Only owner can call this function");
		require(!paused, "The contract is paused");
		_to.transfer(address(this).balance);
	}

	function destroySmartContract(address payable _to) public {
		require(msg.sender == owner, "Only owner can call this function");
		require(!paused, "The contract is paused");
		selfdestruct(_to);
	}
}