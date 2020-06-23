/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package triones

import (
	"fmt"
	"github.com/ontio/celo-ontid/common"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/triones-node-tool/config"
)

func RegisterTriones(ontSdk *sdk.OntologySdk) bool {
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = registerCandidate(ontSdk, user, config.DefConfig.PeerPublicKey, config.DefConfig.InitPos)
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

func QuitTriones(ontSdk *sdk.OntologySdk) bool {
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = quitNode(ontSdk, user, config.DefConfig.PeerPublicKey)
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

func WithdrawInitPos(ontSdk *sdk.OntologySdk) bool {
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = withdraw(ontSdk, user, []string{config.DefConfig.PeerPublicKey}, []uint32{config.DefConfig.InitPos})
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

func WithdrawOng(ontSdk *sdk.OntologySdk) bool {
	user, ok := common.GetAccountByPassword(ontSdk, config.DefConfig.WalletPath)
	if !ok {
		return false
	}
	ok = withdrawFee(ontSdk, user)
	if !ok {
		return false
	}
	common.WaitForBlock(ontSdk)
	return true
}

func GetTrionesInfo(ontSdk *sdk.OntologySdk) bool {
	peerPoolMap, err := getPeerPoolMap(ontSdk)
	if err != nil {
		fmt.Println("getPeerPoolMap failed ", err)
		return false
	}

	peerPoolItem, ok := peerPoolMap.PeerPoolMap[config.DefConfig.PeerPublicKey]
	if !ok {
		fmt.Println("Can't find peerPubkey in peerPoolMap")
		return false
	}
	fmt.Println("peerPoolItem.Index is:", peerPoolItem.Index)
	fmt.Println("peerPoolItem.PeerPubkey is:", peerPoolItem.PeerPubkey)
	fmt.Println("peerPoolItem.Address is:", peerPoolItem.Address.ToBase58())
	fmt.Println("peerPoolItem.Status is:", peerPoolItem.Status)
	fmt.Println("peerPoolItem.InitPos is:", peerPoolItem.InitPos)
	fmt.Println("peerPoolItem.TotalPos is:", peerPoolItem.TotalPos)
	return true
}
