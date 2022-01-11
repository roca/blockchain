require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: "0.8.4",
  defaultNetwork: "localhost",
  networks: {
    localhost: {
      url: "http://ganache:8545"
      // accounts: [
      //   '0x8bdced7c7924e9e4e671260b753e1edfaa7f44b89b34869ed32a005d30d06c94',
      //   '0x4e21f2cbea383add69800135c06fd37136864e3437d818296d57939df7e404fc',
      //   '0x75d7601d7ff45d0aa0046bf38322fc1ce1170b477950bd4b3b0352639409119f',
      //   '0x5aa96a6c47bc135a04ecf6416b5ca9cbc927c7484a7193ab236ad300e6ad05fa',
      //   '0xa9feef5f24b1119b73450566929a3af5325969fc1f6bba934324804732d196e1',
      //   '0xf9b23a419b4330771ac17c981c59b6abd099ce2815aceec9e8beb844c33b7d65',
      //   '0x09208931dec19b8b801ed0960c140f69aad132155e9850939c909e6147c818d7',
      //   '0xcf32da3b37c18b3d12faaaf4aadfa4474010963d633ea6d57b85cdabb408d967',
      //   '0xe8c8f933d3bd1ba937583139e7639cd21828f6b9898d904e8f467c0c44111ef0',
      //   '0x3f9a3b528be62be8b928bf6eadbf148fcb27c4e54a54ef47ed21f6542eb52e44'
      // ]
    },
    hardhat: {
    },
    rinkeby: {
      url: "https://eth-rinkeby.alchemyapi.io/v2/123abc123abc123abc123abc123abcde"
      //accounts: [privateKey1, privateKey2, ...]
    }
  },
};
