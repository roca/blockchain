// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract SendMoneyExample {

	uint public balanceReceived;

	function receiveMone() public payable {
		balanceReceived += msg.value;
	}

	function getBalance() public view returns (uint) {
		return address(this).balance;
	}

	function withdrawMoney() public{
		address payable to  = payable(msg.sender);

		to.transfer(this.getBalance());
	}

	function withdrawMoneyTo(address payable _to) public{
		_to.transfer(this.getBalance());
	}
}