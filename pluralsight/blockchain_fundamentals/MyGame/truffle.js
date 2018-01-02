module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*" // Match any network id
    },
    production: {
      host: "bclbzx5doosi.eastus.cloudapp.azure.com",
      port: 8545,
      gas: 3000000,
      network_id: "*" // Match any network id
    }
  }
};
