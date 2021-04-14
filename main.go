package main

import (
	"gobitcoin/base"
	"gobitcoin/cli"
)

func main() {
	bc := base.NewBlockchain()
	defer bc.Db.Close()

	c := cli.CLI{bc}
	c.Run()
}
