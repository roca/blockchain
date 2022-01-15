// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract Item {
	uint public priceInWei;
	uint public pricePaid;
	uint public index;

	ItemManager manager;

	constructor(uint _priceInWei, uint _index, ItemManager _manager)  {
		priceInWei = _priceInWei;
		index = _index;
		manager = _manager;
	}

	receive() external payable {
		require(pricePaid == 0, "Item is paid already");
		require(priceInWei == msg.value, "Only full payments allowed");
		(bool success,) = address(manager).call{value: msg.value}(abi.encodeWithSignature("triggerPayment(uint256)", index));
		require(success,"The transaction wasn't successful, canceling");
		pricePaid = msg.value;
	}

	fallback() external{}
}