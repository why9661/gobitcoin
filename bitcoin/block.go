package bitcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp    int64
	Transactions []*Transaction //will be replaced by the root hash of a merkle tree
	PreviousHash []byte
	Hash         []byte
	Nonce        int
}

func NewBlock(transactions []*Transaction, preHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: preHash,
		Hash:         []byte{},
		Nonce:        0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

func NewGenesisblock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var serializaedBlock bytes.Buffer
	encoder := gob.NewEncoder(&serializaedBlock)
	encoder.Encode(b)

	return serializaedBlock.Bytes()
}

func DerializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&block)

	return &block
}

func (b *Block) HashOfTx() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
