if (typeof window.ethereum !== 'undefined') {
    console.log('√ metamask');
}
console.log(window.ethereum);
//取得元素 通过 class 方式
const ethereumButton = document.querySelector('.enableEthereumButton');
const showAccount = document.querySelector('.showAccount');

const ethereumButton1 = document.querySelector('.deploycontract');

let accounts = [];

//绑定 click 事件
ethereumButton.addEventListener('click', () => {
    getAccount();
});

//绑定 click 事件
ethereumButton1.addEventListener('click', () => {
    deployContract();
});

async function getAccount() {
    accounts = await ethereum.request({ method: 'eth_requestAccounts' });
    console.log(accounts);
    showAccount.innerHTML = accounts[0];//修改 showAccount 元神的 html 为 账户内容
}

async function deployContract() {
    console.log("deploy start....")
    const abi = [ { "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_tokenId", "type": "uint256" } ], "name": "transfer", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" } ];
    const bytecode = "60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206e71e75bd36b7af41706d1c7f01b33d3ce411d945e500e38742ad9a31265be5064736f6c63430008070033"
    //const myContract = new web3.eth.Contract(abi);

    // Deploy the contract
    //myContract.deploy({data: bytecode}).send({from: web3.eth.defaultAccount, gas: 1000000}).then(function (instance) {
     //   console.log("Contract deployed to: ", instance.options.address);
    //});
    const instance=App.contract.Demo.Deployed();
    intance.set(valstr1,valint2);


}
