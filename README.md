# triones-node-tool

## Instruction

### 1. Clone repo

```shell
git clone https://github.com/ontio/triones-node-tool.git
```

### 2. Build or download latest release

bulid:

```shell
go build main.go
```

or download latest release:

https://github.com/ontio/triones-node-tool/releases

### 3. Update config file

```shell
vim config.json
```

content of config.json：

```json
{
  "JsonRpcAddress":"http://polaris1.ont.io:20336",
  "WalletPath": "wallet.dat",
  "PeerPublicKey": "02d500543eaa4110b8d5f8df4fabe31a8caabaa58cb8512baa8af15a303606efe4",
  "InitPos":10000,
  "GasPrice":500,
  "GasLimit":20000
}
```

`JsonRpcAddress`：rpc of ontology nodes

for mainnet: 
`"http://dappnode1.ont.io:20336","http://dappnode2.ont.io:20336","http://dappnode3.ont.io:20336","http://dappnode4.ont.io:20336",`

for polaris testnet: 
`"http://polaris1.ont.io:20336"，"http://polaris2.ont.io:20336"，"http://polaris3.ont.io:20336"，"http://polaris4.ont.io:20336",`

`WalletPath`: path of wallet file, this wallet is deposit account

`PeerPublicKey`: public key of node's operation account

`InitPos`: init deposit

### 4. Run command line

list of supported command line: 

| command line                | function                                                     |
| --------------------------- | ------------------------------------------------------------ |
| `./main -t RegisterTriones` | register node to triones nodes, it will deposit value of InitPos from deposit account, charge 500 ONG as fee |
| `./main -t QuitTriones`     | quit triones, it will quit node from triones, candidate node will quit next epoch, consensus node will quit next next epoch |
| `./main -t WithdrawInitPos` | withdraw InitPos, it can be called to withdraw InitPos after node quits successfully |
| `./main -t WithdrawOng`     | withdraw node benefits                                       |
| `./main -t GetTrionesInfo`  | query node info, for status: 1 means candidate node, 2 means consensus node, 3 means quited consensus node, 4 means quited candidate node. TotalPos means total deposit from users. |

And now you can run your command and input your password if needed.
