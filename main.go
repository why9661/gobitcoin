package main

import (
	"fmt"
	"goblockchain/block"
)

func main() {
	bc := block.NewBlockChain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")

	for _, b := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", b.PreviousHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Println()
	}
}
