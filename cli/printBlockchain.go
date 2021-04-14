package cli

import (
	"fmt"
	. "gobitcoin/base"
	"strconv"
)

func (c *CLI) printBlockchain() {
	iterator := c.Bc.Iterator()

	for {
		block := iterator.Next()

		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
