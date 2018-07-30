/*
Copyright ArxanFintech Technology Ltd. 2017-2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package wallet

////////////////////////////////////////////////////////////////////////////////
// Transaction Logs Structs

type TransactionLogs map[string]*TransactionLog // key is endpoint

type TransactionLog struct {
	Utxo []*UTXO       `json:"utxo,omitempty"` // unspent transaction output
	Stxo []*SpentTxOUT `json:"stxo,omitempty"` // spent transaction output
}

type UTXO struct {
	// SourceTxDataHash the Bitcoin hash (double sha256) of
	// the given transaction
	SourceTxDataHash string `protobuf:"bytes,1,opt,name=sourceTxDataHash" json:"sourceTxDataHash,omitempty" `
	// Ix index of output array in the transaction
	Ix string `protobuf:"varint,2,opt,name=ix" json:"ix,omitempty" `
	// ColoredToken ID
	CTokenId string `protobuf:"bytes,3,opt,name=cTokenId" json:"cTokenId,omitempty" `
	// ColorType
	CType int32 `protobuf:"varint,4,opt,name=cType" json:"cType,omitempty"`
	// token amount
	Value int64 `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	// who will receive this txout
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty" `
	// until xx timestamp, any one cant spend the txout
	// -1 means no check
	Until int64 `protobuf:"varint,6,opt,name=until" json:"until,omitempty"`
	// script
	Script []byte `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	// CreatedAt
	CreatedAt *Timestamp `protobuf:"bytes,8,opt,name=createdAt" json:"createdAt,omitempty"`
	// Founder who created this tx
	Founder string `protobuf:"bytes,9,opt,name=founder" json:"founder,omitempty" `
	TxType  string `protobuf:"varint,10,opt,name=txType" json:"txType,omitempty"`
	// BCTxID blockchain transaction id
	BCTxID string `protobuf:"bytes,11,opt,name=bcTxID" json:"bcTxID,omitempty"`
}

// SpentTxOUT
type SpentTxOUT struct {
	// SourceTxDataHash the Bitcoin hash (double sha256) of
	// the given transaction
	SourceTxDataHash string `protobuf:"bytes,1,opt,name=sourceTxDataHash" json:"sourceTxDataHash,omitempty" `
	// Ix index of output array in the transaction
	Ix string `protobuf:"varint,2,opt,name=ix" json:"ix,omitempty" `
	// ColoredToken ID
	CTokenId string `protobuf:"bytes,3,opt,name=cTokenId" json:"cTokenId,omitempty" `
	// ColorType
	CType int32 `protobuf:"varint,4,opt,name=cType" json:"cType,omitempty"`
	// token amount
	Value int64 `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	// who will receive this txout
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty" `
	// until xx timestamp, any one cant spend the txout
	// -1 means no check
	Until int64 `protobuf:"varint,6,opt,name=until" json:"until,omitempty"`
	// script
	Script []byte `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	// CreatedAt
	CreatedAt *Timestamp `protobuf:"bytes,8,opt,name=createdAt" json:"createdAt,omitempty"`
	// SpentTxDataHash
	SpentTxDataHash string `protobuf:"bytes,9,opt,name=spentTxDataHash" json:"spentTxDataHash,omitempty" `
	// SpentAt ...
	SpentAt *Timestamp `protobuf:"bytes,10,opt,name=spentAt" json:"spentAt,omitempty"`
	// Founder who created this tx
	Founder string `protobuf:"bytes,11,opt,name=founder" json:"founder,omitempty"`
	TxType  string `protobuf:"varint,12,opt,name=txType" json:"txType,omitempty"`
	// BCTxID blockchain transaction id
	BCTxID string `protobuf:"bytes,13,opt,name=bcTxID" json:"bcTxID,omitempty"`
}

// Timestamp Structure
type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}
