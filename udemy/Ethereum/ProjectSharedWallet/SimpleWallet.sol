// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

//import "./Owned.sol";

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol";

contract Allowance is Ownable {
	event AllowanceChanged(address indexed _forWho, address indexed _fromWhom, uint _oldAmount , uint _newAmount);

	mapping(address => uint) public allowance;

	function addAllowance(address _who, uint _amount) public onlyOwner{
		emit AllowanceChanged(_who, msg.sender, allowance[_who], _amount);
		allowance[_who] = _amount;
	}


    function isOwner() public view returns(bool) { 
		return owner() == msg.sender;
	}

	function getBalance() public view returns (uint) {
		return address(this).balance;
	}

	modifier ownerOrAllowed(uint _amount) {
		require( isOwner() || allowance[msg.sender] >= _amount, "You ar not allowed");
		_;
	}

	function reduceAllowance(address _who, uint _amount) internal {
		emit AllowanceChanged(_who, msg.sender, allowance[_who],  allowance[_who] - _amount);
		allowance[_who] -= _amount;
	}
}

contract SimpleWallet is Allowance {

	event MoneySent(address indexed _beneficiary, uint _amount);
	event MoneyReceived(address indexed _from, uint _amount);

	function withdrawMoney(address payable _to, uint _amount) public ownerOrAllowed(_amount) {
		require(_amount <= address(this).balance, "There are not enough funds stored in the smart contract");
		if (!isOwner()) {
			reduceAllowance(msg.sender,_amount);
		}
		emit MoneySent(_to, _amount);
		_to.transfer(_amount);
	}

	receive() external payable {
		emit MoneyReceived(msg.sender, msg.value);
	}
}