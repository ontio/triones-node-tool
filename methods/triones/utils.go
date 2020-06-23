package triones

import (
	"bytes"

	log4 "github.com/alecthomas/log4go"
	"github.com/ontio/ont-relayer/common"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-tool/config"
	ontcommon "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/errors"
	"github.com/ontio/ontology/smartcontract/service/native/governance"
	"github.com/ontio/ontology/smartcontract/service/native/utils"
)

var OntIDVersion = byte(0)

func registerCandidate(ontSdk *sdk.OntologySdk, user *sdk.Account, peerPubkey string, initPos uint32) bool {
	params := &governance.RegisterCandidateParam{
		PeerPubkey: peerPubkey,
		Address:    user.Address,
		InitPos:    initPos,
	}
	method := "registerCandidate"
	contractAddress := utils.GovernanceContractAddress
	tx, err := ontSdk.Native.NewNativeInvokeTransaction(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		log4.Error("NewNativeInvokeTransaction error :", err)
		return false
	}
	err = ontSdk.SignToTransaction(tx, user)
	if err != nil {
		log4.Error("SignToTransaction error :", err)
		return false
	}
	txHash, err := ontSdk.SendTransaction(tx)
	if err != nil {
		log4.Error("SendTransaction error :", err)
		return false
	}
	log4.Info("registerCandidate txHash is :", txHash.ToHexString())
	return true
}

func quitNode(ontSdk *sdk.OntologySdk, user *sdk.Account, peerPubkey string) bool {
	params := &governance.QuitNodeParam{
		PeerPubkey: peerPubkey,
		Address:    user.Address,
	}
	contractAddress := utils.GovernanceContractAddress
	method := "quitNode"
	txHash, err := ontSdk.Native.InvokeNativeContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		user, user, OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		log4.Error("invokeNativeContract error :", err)
		return false
	}
	log4.Info("quitNode txHash is :", txHash.ToHexString())
	return true
}

func withdraw(ontSdk *sdk.OntologySdk, user *sdk.Account, peerPubkeyList []string, withdrawList []uint32) bool {
	params := &governance.WithdrawParam{
		Address:        user.Address,
		PeerPubkeyList: peerPubkeyList,
		WithdrawList:   withdrawList,
	}
	contractAddress := utils.GovernanceContractAddress
	method := "withdraw"
	txHash, err := ontSdk.Native.InvokeNativeContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		user, user, OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		log4.Error("invokeNativeContract error :", err)
		return false
	}
	log4.Info("withdraw txHash is :", txHash.ToHexString())
	return true
}

func withdrawFee(ontSdk *sdk.OntologySdk, user *sdk.Account) bool {
	params := &governance.WithdrawFeeParam{
		Address: user.Address,
	}
	contractAddress := utils.GovernanceContractAddress
	method := "withdrawFee"
	txHash, err := ontSdk.Native.InvokeNativeContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		user, user, OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		log4.Error("invokeNativeContract error :", err)
		return false
	}
	log4.Info("withdrawFee txHash is :", txHash.ToHexString())
	return true
}

func getGovernanceView(ontSdk *sdk.OntologySdk) (*governance.GovernanceView, error) {
	contractAddress := utils.GovernanceContractAddress
	governanceView := new(governance.GovernanceView)
	key := []byte(governance.GOVERNANCE_VIEW)
	value, err := ontSdk.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "getStorage error")
	}
	if err := governanceView.Deserialize(bytes.NewBuffer(value)); err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "deserialize, deserialize governanceView error!")
	}
	return governanceView, nil
}

func getView(ontSdk *sdk.OntologySdk) (uint32, error) {
	governanceView, err := getGovernanceView(ontSdk)
	if err != nil {
		return 0, errors.NewDetailErr(err, errors.ErrNoCode, "getGovernanceView error")
	}
	return governanceView.View, nil
}

func getPeerPoolMap(ontSdk *sdk.OntologySdk) (*governance.PeerPoolMap, error) {
	contractAddress := utils.GovernanceContractAddress
	view, err := getView(ontSdk)
	if err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "getView error")
	}
	peerPoolMap := &governance.PeerPoolMap{
		PeerPoolMap: make(map[string]*governance.PeerPoolItem),
	}
	viewBytes := governance.GetUint32Bytes(view)
	key := common.ConcatKey([]byte(governance.PEER_POOL), viewBytes)
	value, err := ontSdk.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "getStorage error")
	}
	if err := peerPoolMap.Deserialization(ontcommon.NewZeroCopySource(value)); err != nil {
		return nil, errors.NewDetailErr(err, errors.ErrNoCode, "deserialize, deserialize peerPoolMap error!")
	}
	return peerPoolMap, nil
}
