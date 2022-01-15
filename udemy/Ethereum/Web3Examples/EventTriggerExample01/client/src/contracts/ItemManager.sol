// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract ItemManager is Ownable {

	enum SpupplyChainState {Created, Paid, Delivered}

	struct S_Item {
		Item item;
		string _identifier;
		uint _itemPrice;
		SpupplyChainState _state;
	}

	mapping(uint => S_Item) public items;
	uint itemIndex;

	event SupplyChainStep(uint _itemIndex, uint _step, address _itemAddress);

	function createItem(string memory _identifier, uint _itemPrice) public onlyOwner returns(address)  {
		Item newItem = new Item(_itemPrice,itemIndex,this);
		S_Item memory item = S_Item(newItem,_identifier, _itemPrice, SpupplyChainState.Created);
		items[itemIndex] = item;
		address itemAddress = address(items[itemIndex].item);
		emit SupplyChainStep(itemIndex, uint(SpupplyChainState.Created),itemAddress);
		itemIndex++;
		return itemAddress;
	}

	function triggerPayment(uint _itemIndex) public payable {
		require(items[_itemIndex]._itemPrice == msg.value,"Only full payments accepted");
		require(items[_itemIndex]._state == SpupplyChainState.Created, "Item is further in the chain");

		items[_itemIndex]._state = SpupplyChainState.Paid;
		emit SupplyChainStep(_itemIndex, uint(SpupplyChainState.Paid),address(items[_itemIndex].item));
	}

	function triggerDelivery(uint _itemIndex) public onlyOwner {
		require(items[_itemIndex]._state == SpupplyChainState.Paid, "Item is further in the chain");

		items[_itemIndex]._state = SpupplyChainState.Delivered;
		emit SupplyChainStep(_itemIndex, uint(SpupplyChainState.Delivered),address(items[_itemIndex].item));
	}
}