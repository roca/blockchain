const ItemManger = artifacts.require("./ItemManager.sol");

contract("ItemManager", accounts => {
	it("SHould be able to add an item", async () => {
		const itemManagerInstance = await ItemManger.deployed();
		const itemName = "example_1";
		const itemPrice = 100;

		const result = await itemManagerInstance.createItem(itemName, itemPrice, { from: accounts[0] });
		assert.equal(result.logs[0].args._itemIndex, 0, "It's not the first item");

		const item = await itemManagerInstance.items(0);
		assert.equal(item._identifier, itemName, "Item name is not correct");
	});
});