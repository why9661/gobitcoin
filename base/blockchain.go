package base

import bolt "go.etcd.io/bbolt"

const (
	dbFile = "blockchain.db"
	bucket = "blocks"
)

type Blockchain struct {
	tip []byte //the last block's hash
	Db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	_ = bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	_ = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		_ = b.Put(newBlock.Hash, newBlock.Serialize())
		_ = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})

}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)

	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, _ := tx.CreateBucket([]byte(bucket))
			_ = b.Put(genesis.Hash, genesis.Serialize())
			_ = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.Db}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	_ = i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		serializedBlock := b.Get(i.currentHash)
		block = Deserialize(serializedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash

	return block
}
