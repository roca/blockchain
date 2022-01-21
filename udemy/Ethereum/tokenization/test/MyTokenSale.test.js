const MyTokenSale = artifacts.require("./MyTokenSale");
const MyToken = artifacts.require("./MyToken");

const {chai, BN} = require("./setupchai");
const expect = chai.expect;

require('dotenv').config({ path: '../.env' });

contract("MyTokenSale Test", accounts => {
	const [deployerAccount, recipient1, recipient2] = accounts;

	beforeEach(async() =>{
		this.myTokenSale = await MyTokenSale.deployed();
		this.myToken = await MyToken.deployed();
	});

	it("should not have any tokens in my deployerAccount", async () => {
		let instance = this.myToken;
		let balance = instance.balanceOf(deployerAccount);

		return expect(balance).to.eventually.be.a.bignumber.equal(new BN(0));
	});

	it("all tokens should be in the myTokenSale smart contract by default", async () => {
		let instance = this.myToken;
		let balanceOfTokenSaleSmartContract = await instance.balanceOf(MyTokenSale.address);
		let totalSupply = await instance.totalSupply();

		return expect(balanceOfTokenSaleSmartContract).to.be.a.bignumber.equal(totalSupply);
	});

	it("should be possible to buy tokens", async () => {
		let tokenInstance = this.myToken;
		let tokenSaleInstance = this.myTokenSale;
		let balanceBefore = await tokenInstance.balanceOf(deployerAccount);

		await expect(tokenSaleInstance.sendTransaction({from: deployerAccount, value: web3.utils.toWei("1", "wei")})).to.be.fulfilled;
		return expect(tokenInstance.balanceOf(deployerAccount)).to.eventually.be.a.bignumber.equal(balanceBefore.add(new BN(1)));
	});


});