package command

import (
	"strings"
)

type InsertCommand struct {
	Meta
}

func (c *InsertCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *InsertCommand) Synopsis() string {
	return ""
}

func (c *InsertCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
