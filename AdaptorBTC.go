/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018
 */
package btcadaptor

import (
	"errors"

	"github.com/palletone/adaptor"
)

type RPCParams struct {
	Host      string `json:"host"`
	RPCUser   string `json:"rpcUser"`
	RPCPasswd string `json:"rpcPasswd"`
	CertPath  string `json:"certPath"`
}

type AdaptorBTC struct {
	NetID int
	RPCParams
}

const (
	NETID_MAIN = iota
	NETID_TEST
)

/*IUtility*/
//创建一个新的私钥
func (abtc *AdaptorBTC) NewPrivateKey(input *adaptor.NewPrivateKeyInput) (*adaptor.NewPrivateKeyOutput, error) {
	prikey, err := NewPrivateKey(abtc.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.NewPrivateKeyOutput{PrivateKey: prikey}
	return &result, nil
}

//根据私钥创建公钥
func (abtc *AdaptorBTC) GetPublicKey(input *adaptor.GetPublicKeyInput) (*adaptor.GetPublicKeyOutput, error) {
	pubkey, err := GetPublicKey(input.PrivateKey, abtc.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetPublicKeyOutput{PublicKey: pubkey}
	return &result, nil
}

//根据Key创建地址
func (abtc *AdaptorBTC) GetAddress(key *adaptor.GetAddressInput) (*adaptor.GetAddressOutput, error) {
	addr, err := PubKeyToAddress(key.Key, abtc.NetID)
	if err != nil {
		return nil, err
	}
	result := adaptor.GetAddressOutput{Address: addr}
	return &result, nil
}

//获得原链的地址和PalletOne的地址的映射 //btc， not implement
func (abtc *AdaptorBTC) GetPalletOneMappingAddress(addr *adaptor.GetPalletOneMappingAddressInput) (*adaptor.GetPalletOneMappingAddressOutput, error) { //todo
	return nil, errors.New("todo")
}

//对一条消息进行签名
func (abtc *AdaptorBTC) SignMessage(input *adaptor.SignMessageInput) (*adaptor.SignMessageOutput, error) {
	return SignMessage(input)
}

//对签名进行验证
func (abtc *AdaptorBTC) VerifySignature(input *adaptor.VerifySignatureInput) (*adaptor.VerifySignatureOutput, error) {
	return VerifySignature(input)
}

//对一条交易进行签名，并返回签名结果
func (abtc *AdaptorBTC) SignTransaction(input *adaptor.SignTransactionInput) (*adaptor.SignTransactionOutput, error) { //todo
	return nil, errors.New("todo")
}

//将未签名的原始交易与签名进行绑定，返回一个签名后的交易
func (abtc *AdaptorBTC) BindTxAndSignature(input *adaptor.BindTxAndSignatureInput) (*adaptor.BindTxAndSignatureOutput, error) { //todo
	return nil, errors.New("todo")
}

//根据交易内容，计算交易Hash
func (abtc *AdaptorBTC) CalcTxHash(input *adaptor.CalcTxHashInput) (*adaptor.CalcTxHashOutput, error) { //todo
	return nil, errors.New("todo")
}

//将签名后的交易广播到网络中,如果发送交易需要手续费，指定最多支付的手续费
func (abtc *AdaptorBTC) SendTransaction(input *adaptor.SendTransactionInput) (*adaptor.SendTransactionOutput, error) { //todo zxl
	return nil, errors.New("todo")
}

//根据交易ID获得交易的基本信息
func (abtc *AdaptorBTC) GetTxBasicInfo(input *adaptor.GetTxBasicInfoInput) (*adaptor.GetTxBasicInfoOutput, error) { //todo zxl
	return nil, errors.New("todo")
}

//查询获得一个区块的信息
func (abtc *AdaptorBTC) GetBlockInfo(input *adaptor.GetBlockInfoInput) (*adaptor.GetBlockInfoOutput, error) { //todo zxl
	return nil, errors.New("todo")
}

/*ICryptoCurrency*/
//获取某地址下持有某资产的数量,返回数量为该资产的最小单位
func (abtc *AdaptorBTC) GetBalance(input *adaptor.GetBalanceInput) (*adaptor.GetBalanceOutput, error) { //todo zxl
	return nil, errors.New("todo")
}

//获取某资产的小数点位数
func (abtc *AdaptorBTC) GetAssetDecimal(asset *adaptor.GetAssetDecimalInput) (*adaptor.GetAssetDecimalOutput, error) {
	result := adaptor.GetAssetDecimalOutput{Decimal: 8}
	return &result, nil
}

//创建一个转账交易，但是未签名
func (abtc *AdaptorBTC) CreateTransferTokenTx(input *adaptor.CreateTransferTokenTxInput) (*adaptor.CreateTransferTokenTxOutput, error) { //todo
	return nil, errors.New("todo")
}

//获取某个地址对某种Token的交易历史,支持分页和升序降序排列
func (abtc *AdaptorBTC) GetAddrTxHistory(input *adaptor.GetAddrTxHistoryInput) (*adaptor.GetAddrTxHistoryOutput, error) { //todo
	return nil, errors.New("todo")
}

//根据交易ID获得对应的转账交易
func (abtc *AdaptorBTC) GetTransferTx(input *adaptor.GetTransferTxInput) (*adaptor.GetTransferTxOutput, error) { //todo zxl
	return nil, errors.New("todo")
}

//创建一个多签地址，该地址必须要满足signCount个签名才能解锁
func (abtc *AdaptorBTC) CreateMultiSigAddress(input *adaptor.CreateMultiSigAddressInput) (*adaptor.CreateMultiSigAddressOutput, error) {
	return CreateMultiSigAddress(input, abtc.NetID)
}

//func (abtc AdaptorBTC) NewPrivateKey() (wifPriKey string) {
//	return NewPrivateKey(abtc.NetID)
//}
//func (abtc AdaptorBTC) GetPublicKey(wifPriKey string) (pubKey string) {
//	return GetPublicKey(wifPriKey, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetAddress(wifPriKey string) (address string) {
//	return GetAddress(wifPriKey, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetAddressByPubkey(pubKeyHex string) (string, error) {
//	return GetAddressByPubkey(pubKeyHex, abtc.NetID)
//}
//func (abtc AdaptorBTC) CreateMultiSigAddress(params *adaptor.CreateMultiSigParams) (*adaptor.CreateMultiSigResult, error) {
//	return CreateMultiSigAddress(params, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) GetUTXO(params *adaptor.GetUTXOParams) (*adaptor.GetUTXOResult, error) {
//	return GetUTXO(params, &abtc.RPCParams, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetUTXOHttp(params *adaptor.GetUTXOHttpParams) (*adaptor.GetUTXOHttpResult, error) {
//	return GetUTXOHttp(params, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) RawTransactionGen(params *adaptor.RawTransactionGenParams) (*adaptor.RawTransactionGenResult, error) {
//	return RawTransactionGen(params, abtc.NetID)
//}
//func (abtc AdaptorBTC) DecodeRawTransaction(params *adaptor.DecodeRawTransactionParams) (*adaptor.DecodeRawTransactionResult, error) {
//	return DecodeRawTransaction(params, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetTransactionByHash(params *adaptor.GetTransactionByHashParams) (*adaptor.GetTransactionByHashResult, error) {
//	return GetTransactionByHash(params, &abtc.RPCParams)
//}
//func (abtc AdaptorBTC) GetTransactionHttp(params *adaptor.GetTransactionHttpParams) (*adaptor.GetTransactionHttpResult, error) {
//	return GetTransactionHttp(params, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) SignTransaction(params *adaptor.SignTransactionParams) (*adaptor.SignTransactionResult, error) {
//	return SignTransaction(params, abtc.NetID)
//}
//func (abtc AdaptorBTC) SignTxSend(params *adaptor.SignTxSendParams) (*adaptor.SignTxSendResult, error) {
//	return SignTxSend(params, &abtc.RPCParams, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetBalance(params *adaptor.GetBalanceParams) (*adaptor.GetBalanceResult, error) {
//	return GetBalance(params, &abtc.RPCParams, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetBalanceHttp(params *adaptor.GetBalanceHttpParams) (*adaptor.GetBalanceHttpResult, error) {
//	return GetBalanceHttp(params, abtc.NetID)
//}
//func (abtc AdaptorBTC) GetTransactions(params *adaptor.GetTransactionsParams) (*adaptor.TransactionsResult, error) {
//	return GetTransactions(params, &abtc.RPCParams, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) SendTransaction(params *adaptor.SendTransactionParams) (*adaptor.SendTransactionResult, error) {
//	return SendTransaction(params, &abtc.RPCParams)
//}
//func (abtc AdaptorBTC) SendTransactionHttp(params *adaptor.SendTransactionHttpParams) (*adaptor.SendTransactionHttpResult, error) {
//	return SendTransactionHttp(params, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) MergeTransaction(params *adaptor.MergeTransactionParams) (*adaptor.MergeTransactionResult, error) {
//	return MergeTransaction(params, abtc.NetID)
//}
//
//func (abtc AdaptorBTC) SignMessage(params *adaptor.SignMessageParams) (*adaptor.SignMessageResult, error) {
//	return SignMessage(params)
//}
//
//func (abtc AdaptorBTC) VerifyMessage(params *adaptor.VerifyMessageParams) (*adaptor.VerifyMessageResult, error) {
//	return VerifyMessage(params, abtc.NetID)
//}
