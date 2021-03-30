package main

import (
	"goblockchain/block"
	"goblockchain/cmd"
)

func main() {
	bc := block.NewBlockchain()
	defer bc.Db.Close()

	cli := cmd.CLI{bc}
	cli.Run()
}
