package bitcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 50

type Transaction struct {
	ID   []byte //hash of transaction
	VIn  []TxInput
	VOut []TxOutput
}

type TxOutput struct {
	Value        int
	ScriptPubKey string
}

type TxInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func (tx *Transaction) SetID() {
	var encodedTx bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encodedTx)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encodedTx.Bytes())
	tx.ID = hash[:]
}

//create coinbase transaction
func NewCoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{subsidy, to}
	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}
