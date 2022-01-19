const MyToken = artifacts.require("./MyToken");

var chai = require("chai");
const BN = web3.utils.BN;
const chaiBN = require("chai-bn")(BN);
chai.use(chaiBN);

var chaiAsPromised = require("chai-as-promised");
chai.use(chaiAsPromised);

const expect = chai.expect;

contract("MyToken Test", accounts => {
	const [deployerAccount, recipient1, recipient2] = accounts;

	it("all tokens should be in my account", async () => {
		let instance = await MyToken.deployed();
		let totalSupply = await instance.totalSupply();
		let balance = instance.balanceOf(deployerAccount);

		expect(balance).to.eventually.be.a.bignumber.equal(totalSupply);
	});

	it("should be possible to send tokens between accounts", async () => {
		const sendTokens = 1;
		let instance = await MyToken.deployed();
		let totalSupply = await instance.totalSupply();
		let balance = instance.balanceOf(deployerAccount);

		expect(balance).to.eventually.be.a.bignumber.equal(totalSupply);
		expect(instance.transfer(recipient1, sendTokens)).to.eventually.be.fulfilled;
		expect(instance.balanceOf(deployerAccount)).to.eventually.be.a.bignumber.equal(totalSupply.sub(new BN(sendTokens)));
		expect(instance.balanceOf(recipient1)).to.eventually.be.a.bignumber.equal(new BN(sendTokens));
	});

	// it("should not be possible to send more tokens than available in total", async () => {
		
	// 	let instance = await MyToken.deployed();
		
	// 	let balanceOfDeployer = await instance.balanceOf(deployerAccount);

	// 	expect(instance.transfer(recipient1,new BN(balanceOfDeployer + 1)).to.eventually.be.rejected);
	// 	expect(instance.balanceOf(deployerAccount)).to.eventually.be.a.bignumber.equal(balanceOfDeployer);
		
	// });
});