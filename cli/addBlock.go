package cli

import (
	"fmt"
)

func (c *CLI) addBlock(data string) {
	c.Bc.AddBlock(data)
	fmt.Println("Success")
}
