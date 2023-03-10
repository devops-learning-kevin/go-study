require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */

const PrivateKey = "8f832ddb358653e315aaa3677726f913660cb6bc3d63982f46a94ae14084f47b";
module.exports = {
  solidity: "0.8.18",
  paths: {
    artifacts: './src/artifacts',
  },
  defaultNetwork: "huaweiyun",
  networks: {
    hardhat:{
      chainId: 67,
    },
    huaweiyun: {
      chainId: 67,
      url: "http://110.41.154.234:8545/",
      accounts: ["8f832ddb358653e315aaa3677726f913660cb6bc3d63982f46a94ae14084f47b"],
    },

    polygon: {
      chainId: 137,
      url: "https://rpc.ankr.com/polygon",
      accounts: ["d5593fe8c33cc12a2a42b86a47d45735b561370602e697e3580085eba8035c17"],
    },

    tbsc: {
      chainId: 97,
      url: "https://endpoints.omniatech.io/v1/bsc/testnet/public",
      accounts: ["371f72ee8f10a0670393c8049feb8f765b878b5dd5e307c781aedd484b085a85"],
    }
  }
};
