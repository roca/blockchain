const path = require("path");

require('dotenv').config({ path: './.env' });
const HDWalletProvider = require("@truffle/hdwallet-provider");
const AccountIndex = 0;

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  contracts_build_directory: path.join(__dirname, "client/src/contracts"),
  networks: {
    ganache: {
      host: "ganache",
      port: 8545,
      network_id: "*"
    },
    ganache_local: {
      provider: function() {
        return new HDWalletProvider(process.env.MNEMONIC, "http://ganache:8545", AccountIndex)
      },
      network_id: "*",
      networkCheckTimeout: 100000
    }
  },
  compilers: {
    solc: {
      version: "0.8.1"
    }
  }
};

// truffle migrate --network ganache_local
