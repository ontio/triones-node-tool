示例代码：https://github.com/ontio/ontology-java-sdk/blob/master/src/main/java/com/github/ontio/smartcontract/nativevm/Governance.java#L799

```java
    public String setPeerCost(Account account,String peerPubkey,int peerCost,Account payerAcct,long gaslimit,long gasprice) throws Exception {
        if(account== null || peerPubkey == null || peerPubkey.equals("") ||payerAcct == null){
            throw new SDKException(ErrorCode.ParamErr("parameters should not be null"));
        }
        if(gaslimit < 0 || gasprice < 0){
            throw new SDKException(ErrorCode.ParamErr("gaslimit and gasprice should not be less than 0"));
        }
        if(peerCost <0 || peerCost > 100){
            throw new SDKException(ErrorCode.ParamErr("peerCost is wrong, it should be 0 <= peerCost <= 100"));
        }
        List list = new ArrayList();
        list.add(new Struct().add(peerPubkey,account.getAddressU160(),peerCost));
        byte[] args = NativeBuildParams.createCodeParamsScript(list);
        Transaction tx = sdk.vm().buildNativeParams(new Address(Helper.hexToBytes(contractAddress)),"setPeerCost",args,payerAcct.getAddressU160().toBase58(),gaslimit, gasprice);
        sdk.signTx(tx,new Account[][]{{account}});
        if(!account.equals(payerAcct)){
            sdk.addSign(tx,payerAcct);
        }
        boolean b = sdk.getConnect().sendRawTransaction(tx.toHexString());
        if (b) {
            return tx.hash().toString();
        }
        return null;
    }
```

参数说明：

`account`: 节点的质押钱包账户

`peerPubkey`: 节点的节点公钥

`peerCost`: 数字0-100，表示节点自己所占收益的百分比，其余部分将分给质押用户。

`payerAcct`: 手续费代付账户，将支付手续费，可以与account相同。

`gaslimit`: 默认20000

`gasprice`: 默认500, 2020年7月7日之后请传入2500

构造交易部分：

```java
Transaction tx = sdk.vm().buildNativeParams(new Address(Helper.hexToBytes(contractAddress)),"setPeerCost",args,payerAcct.getAddressU160().toBase58(),gaslimit, gasprice);
```

签名部分：

```java
sdk.signTx(tx,new Account[][]{{account}});
```

交易的序列化和反序列话参考：

https://github.com/ontio/ontology-java-sdk/blob/master/src/main/java/com/github/ontio/core/transaction/Transaction.java#L37

查询是否设置成功参考：

https://github.com/ontio/ontology-java-sdk/blob/master/src/main/java/com/github/ontio/smartcontract/nativevm/Governance.java#L831

