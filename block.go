package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreviousHash []byte
	Hash         []byte
	Nonce        int
}

func NewBlock(data string, preHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         []byte(data),
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

func NewGenesisblock() *Block {
	return NewBlock("Genesis Block", []byte{})
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
