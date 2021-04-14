package base

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevHash,
		[]byte{},
		0,
	}

	proofOfWork := NewProofOfWork(block)

	nonce, hash := proofOfWork.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	//if err != nil {
	//	panic("Serialize block failed")
	//}

	return result.Bytes()
}

func Deserialize(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	_ = decoder.Decode(&block)
	//if err != nil {
	//	panic("Deserialize block failed")
	//}
	return &block
}
