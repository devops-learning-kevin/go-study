// 引入ethers
const { ethers } = require("ethers");

// 设置JsonRpc连接

const INFURA_ID = "";
const provider = new ethers.providers.JsonRpcProvider(
    `http://110.41.154.234:8545/`
);

const address = "0x1B9d163623a8059794FBb662818fD4249beF975C";
//这一步需要异步操作，因为它很慢，用async await包裹
const main = async () => {
    const balance = await provider.getBalance(address);
    console.log(
        // 需要格式化地址formatEther
        `\nETH Balance of ${address} --> ${ethers.utils.formatEther(balance)} ETH\n`
    );
};

main();

