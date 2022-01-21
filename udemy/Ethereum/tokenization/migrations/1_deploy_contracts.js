var MyToken = artifacts.require("./MyToken");
var MyTokenSale = artifacts.require("./MyTokenSale");
var MyKycContract = artifacts.require("./KycContract");

require('dotenv').config({ path: '../.env' });
const BN = web3.utils.BN;

module.exports = async function(deployer) {
  let addr = await web3.eth.getAccounts();
  let initial_token_supply = new BN(process.env.INITIAL_TOKEN_SUPPLY);
  
  await deployer.deploy(MyToken, process.env.INITIAL_TOKEN_SUPPLY );
  await deployer.deploy(MyKycContract);
  await deployer.deploy(MyTokenSale,1,addr[0],MyToken.address, MyKycContract.address);
  let instance = await MyToken.deployed();
  await instance.transfer(MyTokenSale.address, process.env.INITIAL_TOKEN_SUPPLY );
};
