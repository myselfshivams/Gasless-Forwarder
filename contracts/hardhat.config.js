require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
    solidity: "0.8.20",
    networks: {
        goerli: {
            url: process.env.INFURA_URL,
            accounts: [`0x${process.env.DEPLOYER_PRIVATE_KEY}`]
        },
        sepolia: {
            url: process.env.SEPOLIA_URL,
            accounts: [`0x${process.env.DEPLOYER_PRIVATE_KEY}`]
        }
    },
    etherscan: {
        apiKey: process.env.ETHERSCAN_API_KEY
    }
};
