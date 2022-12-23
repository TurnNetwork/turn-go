// Copyright 2021 The Bubble Network Authors
// This file is part of the bubble library.
//
// The bubble library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The bubble library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the bubble library. If not, see <http://www.gnu.org/licenses/>.

package byteutil

import (
	"encoding/hex"
	"github.com/bubblenet/bubble/x/bubble"
	"math/big"

	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/p2p/enode"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/restricting"
)

var Bytes2X_CMD = map[string]interface{}{
	"string":   BytesToString,
	"*string":  BytesToStringPoint,
	"[8]byte":  BytesTo8Bytes,
	"[16]byte": BytesTo16Bytes,
	"[32]byte": BytesTo32Bytes,
	"[64]byte": BytesTo64Bytes,
	"[]uint8":  BytesToBytes,

	"bool":    BytesToBool,
	"uint8":   BytesToUint8,
	"uint16":  BytesToUint16,
	"*uint16": BytesToUint16Point,
	"uint32":  BytesToUint32,
	"uint64":  BytesToUint64,

	"*big.Int":              BytesToBigInt,
	"[]*big.Int":            BytesToBigIntArr,
	"[][]uint8":             BytesToUint8Arr,
	"enode.IDv0":            BytesToNodeId,
	"[]enode.IDv0":          BytesToNodeIdArr,
	"common.Hash":           BytesToHash,
	"[]common.Hash":         BytesToHashArr,
	"common.Address":        BytesToAddress,
	"*common.Address":       BytesToAddressPoint,
	"[]common.Address":      BytesToAddressArr,
	"common.VersionSign":    BytesToVersionSign,
	"[]common.VersionSign":  BytesToVersionSignArr,
	"bls.PublicKeyHex":      BytesToPublicKeyHex,
	"[]bls.PublicKeyHex":    BytesToPublicKeyHexArr,
	"bls.SchnorrProofHex":   BytesToSchnorrProofHex,
	"[]bls.SchnorrProofHex": BytesToSchnorrProofHexArr,

	"[]restricting.RestrictingPlan": BytesToRestrictingPlanArr,
	"bubble.AccountAsset":           BytesToAccountAsset,
	"bubble.SettlementInfo":         BytesToSettlementInfo,
	"bubble.TxType":                 BytesToBubTxType,
	"bubble.Size":                   BytesToBubSize,
}

func BytesToString(curByte []byte) string {
	//return string(curByte)
	var str string
	if err := rlp.DecodeBytes(curByte, &str); nil != err {
		panic("BytesToString:" + err.Error())
	}
	return str
}

func BytesToStringPoint(curByte []byte) *string {
	if len(curByte) == 0 {
		return nil
	}
	var str string
	if err := rlp.DecodeBytes(curByte, &str); nil != err {
		panic("BytesToString:" + err.Error())
	}
	return &str
}

func BytesTo8Bytes(curByte []byte) [8]byte {
	var arr [8]byte
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesTo8Bytes:" + err.Error())
	}
	return arr
}

func BytesTo16Bytes(curByte []byte) [16]byte {
	var arr [16]byte
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesTo16Bytes:" + err.Error())
	}
	return arr
}

func BytesTo32Bytes(curByte []byte) [32]byte {
	/*var arr [32]byte
	copy(arr[:], curByte)
	return arr*/
	var arr [32]byte
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesTo32Bytes:" + err.Error())
	}
	return arr
}

func BytesTo64Bytes(curByte []byte) [64]byte {
	/*var arr [64]byte
	copy(arr[:], curByte)
	return arr*/
	var arr [64]byte
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesTo64Bytes:" + err.Error())
	}
	return arr
}

func BytesToBytes(curByte []byte) []byte {
	var arr []byte
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesToBytes:" + err.Error())
	}
	return arr
}

func BytesToBool(b []byte) bool {
	var x bool
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToBool:" + err.Error())
	}
	return x
}

func BytesToUint8(b []byte) uint8 {
	var x uint8
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToUint8:" + err.Error())
	}
	return x
}

func BytesToUint16(b []byte) uint16 {
	/*b = append(make([]byte, 2-len(b)), b...)
	return binary.BigEndian.Uint16(b)*/
	var x uint16
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToUint16:" + err.Error())
	}
	return x
}

func BytesToUint16Point(b []byte) *uint16 {
	if len(b) == 0 {
		return nil
	}
	var x *uint16
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToUint16Point:" + err.Error())
	}
	return x
}

func BytesToUint32(b []byte) uint32 {
	/*b = append(make([]byte, 4-len(b)), b...)
	return binary.BigEndian.Uint32(b)*/
	var x uint32
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToUint32:" + err.Error())
	}
	return x
}

func BytesToUint64(b []byte) uint64 {
	/*b = append(make([]byte, 8-len(b)), b...)
	return binary.BigEndian.Uint64(b)*/
	var x uint64
	if err := rlp.DecodeBytes(b, &x); nil != err {
		panic("BytesToUint64:" + err.Error())
	}
	return x
}

func BytesToBigInt(curByte []byte) *big.Int {
	//return new(big.Int).SetBytes(curByte)
	var bigInt *big.Int
	if err := rlp.DecodeBytes(curByte, &bigInt); nil != err {
		panic("BytesToBigInt:" + err.Error())
	}
	return bigInt
}

func BytesToBigIntArr(curByte []byte) []*big.Int {
	var arr []*big.Int
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesToBigIntArr:" + err.Error())
	}
	return arr
}

func BytesToUint8Arr(curByte []byte) [][]uint8 {
	var arr [][]uint8
	if err := rlp.DecodeBytes(curByte, &arr); nil != err {
		panic("BytesToBytesArr:" + err.Error())
	}
	return arr
}

func BytesToNodeId(curByte []byte) enode.IDv0 {
	var nodeId enode.IDv0
	if err := rlp.DecodeBytes(curByte, &nodeId); nil != err {
		panic("BytesToNodeId:" + err.Error())
	}
	return nodeId
}

func BytesToNodeIdArr(curByte []byte) []enode.IDv0 {
	var nodeIdArr []enode.IDv0
	if err := rlp.DecodeBytes(curByte, &nodeIdArr); nil != err {
		panic("BytesToNodeIdArr:" + err.Error())
	}
	return nodeIdArr
}

func BytesToHash(curByte []byte) common.Hash {
	//str := BytesToString(curByte)
	//return common.HexToHash(str)
	var hash common.Hash
	if err := rlp.DecodeBytes(curByte, &hash); nil != err {
		panic("BytesToHash:" + err.Error())
	}
	return hash
}

func BytesToHashArr(curByte []byte) []common.Hash {
	/*str := BytesToString(curByte)
	strArr := strings.Split(str, ":")
	var AHash []common.Hash
	for i := 0; i < len(strArr); i++ {
		AHash = append(AHash, common.HexToHash(strArr[i]))
	}
	return AHash*/

	var hashArr []common.Hash
	if err := rlp.DecodeBytes(curByte, &hashArr); nil != err {
		panic("BytesToHashArr:" + err.Error())
	}
	return hashArr
}

func BytesToAddress(curByte []byte) common.Address {
	//str := BytesToString(curByte)
	//return common.HexToAddress(str)
	var addr common.Address
	if err := rlp.DecodeBytes(curByte, &addr); nil != err {
		panic("BytesToAddress:" + err.Error())
	}
	return addr
}

func BytesToAddressPoint(curByte []byte) *common.Address {
	//str := BytesToString(curByte)
	//return common.HexToAddress(str)
	if len(curByte) == 0 {
		return nil
	}
	var addr common.Address
	if err := rlp.DecodeBytes(curByte, &addr); nil != err {
		panic("BytesToAddress:" + err.Error())
	}
	return &addr
}

func BytesToAddressArr(curByte []byte) []common.Address {
	//str := BytesToString(curByte)
	//return common.HexToAddress(str)
	var addrArr []common.Address
	if err := rlp.DecodeBytes(curByte, &addrArr); nil != err {
		panic("BytesToAddressArr:" + err.Error())
	}
	return addrArr
}

func BytesToVersionSign(currByte []byte) common.VersionSign {
	var version common.VersionSign
	if err := rlp.DecodeBytes(currByte, &version); nil != err {
		panic("BytesToVersionSign:" + err.Error())
	}
	return version
}

func BytesToVersionSignArr(currByte []byte) []common.VersionSign {
	var arr []common.VersionSign
	if err := rlp.DecodeBytes(currByte, &arr); nil != err {
		panic("BytesToVersionSignArr:" + err.Error())
	}
	return arr
}

func BytesToPublicKeyHex(currByte []byte) bls.PublicKeyHex {
	var pub bls.PublicKeyHex
	if err := rlp.DecodeBytes(currByte, &pub); nil != err {
		panic("BytesToPublicKeyHex:" + err.Error())
	}
	return pub
}

func BytesToPublicKeyHexArr(currByte []byte) []bls.PublicKeyHex {
	var arr []bls.PublicKeyHex
	if err := rlp.DecodeBytes(currByte, &arr); nil != err {
		panic("BytesToPublicKeyHexArr:" + err.Error())
	}
	return arr
}

func BytesToSchnorrProofHex(currByte []byte) bls.SchnorrProofHex {
	var proof bls.SchnorrProofHex
	if err := rlp.DecodeBytes(currByte, &proof); nil != err {
		panic("BytesToSchnorrProofHex:" + err.Error())
	}
	return proof
}

func BytesToSchnorrProofHexArr(currByte []byte) []bls.SchnorrProofHex {
	var arr []bls.SchnorrProofHex
	if err := rlp.DecodeBytes(currByte, &arr); nil != err {
		panic("BytesToSchnorrProofHexArr:" + err.Error())
	}
	return arr
}

func BytesToRestrictingPlanArr(curByte []byte) []restricting.RestrictingPlan {
	var planArr []restricting.RestrictingPlan
	if err := rlp.DecodeBytes(curByte, &planArr); nil != err {
		panic("BytesToAddressArr:" + err.Error())
	}
	return planArr
}

func BytesToAccountAsset(curByte []byte) bubble.AccountAsset {
	var accAsset bubble.AccountAsset
	if err := rlp.DecodeBytes(curByte, &accAsset); nil != err {
		panic("BytesToAccountAsset:" + err.Error())
	}
	return accAsset
}

func BytesToSettlementInfo(curByte []byte) bubble.SettlementInfo {
	var settleInfo bubble.SettlementInfo
	if err := rlp.DecodeBytes(curByte, &settleInfo); nil != err {
		panic("BytesToSettlementInfo:" + err.Error())
	}
	return settleInfo
}

func BytesToBubTxType(curByte []byte) bubble.TxType {
	var txType bubble.TxType
	if err := rlp.DecodeBytes(curByte, &txType); nil != err {
		panic("BytesToBubTxType:" + err.Error())
	}
	return txType
}

func BytesToBubSize(curByte []byte) bubble.Size {
	var size bubble.Size
	if err := rlp.DecodeBytes(curByte, &size); nil != err {
		panic("BytesToBubSize:" + err.Error())
	}
	return size
}

func PrintNodeID(nodeID enode.IDv0) string {
	return hex.EncodeToString(nodeID.Bytes()[:8])
}

func Concat(s1 []byte, s2 ...byte) []byte {
	r := make([]byte, len(s1)+len(s2))
	copy(r, s1)
	copy(r[len(s1):], s2)
	return r
}
