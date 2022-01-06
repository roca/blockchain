// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

import "./Owned.sol";

contract SimpleWallet is Owned {

	function withdrawMoney(address payable _to, uint _amount) public onlyOwner {
		_to.transfer(_amount);
	}

	receive() external payable {

	}
}