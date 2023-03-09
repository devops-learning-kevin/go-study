
//取得元素 通过 class 方式
const ethereumButton = document.querySelector('.enableEthereumButton');
const showAccount = document.getElementById('showAccount');
const showBalance = document.getElementById('showBalance');
const sendEthButton = document.querySelector('.sendEthButton');
let accounts = [];
let balance = "";


//绑定 click 事件
ethereumButton.addEventListener('click', () => {
    getAccount();
});

ethereum.on('accountsChanged', async function (accounts) {
    showAccount.value = accounts[0];

    balance = await ethereum.request({
        method: 'eth_getBalance',
        params: [
            accounts[0] ,
            'latest'
        ]
    })
    console.log(balance);
    balance1 = parseInt(balance,16)/Math.pow(10,18)
    showBalance.value = balance1


});

async function getAccount() {
    accounts = await ethereum.request({ method: 'eth_requestAccounts' });
    console.log(accounts);
    showAccount.value = accounts[0];

    balance = await ethereum.request({
        method: 'eth_getBalance',
        params: [
            accounts[0] ,
            'latest'
        ]
    })
    console.log(balance);
    balance1 = parseInt(balance,16)/Math.pow(10,18)
    showBalance.value = balance1
}

//发送 eth
sendEthButton.addEventListener('click', () => {
    var toaddress = document.getElementById('toaddress').value;

    var amount = document.getElementById('amount').value;

    var amount1 = dectoHex(amount*Math.pow(10,18))

    ethereum.request({
        method: 'eth_sendTransaction',
        params: [
            {
                from: accounts[0],
                to: toaddress,
                value: amount1,
                gasPrice: 'e8d4a51000',  //1000 Gwei
                gas: '0x33450',    //210000
            },
        ],
    }).then(
        (txHash) => console.log(txHash)
    ).catch(
        (error) => console.error
    );
});


function dectoHex(n) {
    var str = '';
    if (n == 0) return '0';
    while (n != 0) {
        str = getHexString(n % 16) + str;
        n = parseInt(n / 16);
    }
    return str;
}

function getHexString(num) {
    switch (num) {
        case 10: return 'A';
        case 11: return 'B';
        case 12: return 'C';
        case 13: return 'D';
        case 14: return 'E';
        case 15: return 'F';
        default: return num;
    }
}