const MyToken = artifacts.require("./MyToken");

const {chai, BN} = require("./setupchai");
const expect = chai.expect;

require('dotenv').config({ path: '../.env' });

contract("MyToken Test", accounts => {
	const [deployerAccount, recipient1, recipient2] = accounts;

	beforeEach(async() =>{
		this.myToken = await MyToken.new(new BN(process.env.INITIAL_TOKEN_SUPPLY));
	});

	it("all tokens should be in my account", async () => {
		let instance = this.myToken;
		let totalSupply = await instance.totalSupply();
		let balance = instance.balanceOf(deployerAccount);

		return await expect(balance).to.eventually.be.a.bignumber.equal(totalSupply);
	});

	it("should be possible to send tokens between accounts", async () => {
		const sendTokens = 1;
		let instance = this.myToken;
		let totalSupply = await instance.totalSupply();
		let balance = instance.balanceOf(deployerAccount);

		await expect(balance).to.eventually.be.a.bignumber.equal(totalSupply);
		await expect(instance.transfer(recipient1, sendTokens)).to.eventually.be.fulfilled;
		await expect(instance.balanceOf(deployerAccount)).to.eventually.be.a.bignumber.equal(totalSupply.sub(new BN(sendTokens)));
		return await expect(instance.balanceOf(recipient1)).to.eventually.be.a.bignumber.equal(new BN(sendTokens));
	});

	it("should not be possible to send more tokens than available in total", async () => {
		let instance = this.myToken;
		let balanceOfDeployer = await instance.balanceOf(deployerAccount);

		try {
			await instance.transfer(recipient1,new BN(balanceOfDeployer + 1));
		} catch (error) {
			expect(error.reason).to.include("ERC20: transfer amount exceeds balance");
		}
		//await expect(instance.transfer(recipient1,new BN(balanceOfDeployer + 1)).to.eventually.be.rejected);
		return await expect(instance.balanceOf(deployerAccount)).to.eventually.be.a.bignumber.equal(balanceOfDeployer);
	});
});