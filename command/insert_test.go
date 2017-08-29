package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestInsertCommand_implement(t *testing.T) {
	var _ cli.Command = &InsertCommand{}
}
