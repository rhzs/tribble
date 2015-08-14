package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestStatCommand_implement(t *testing.T) {
	var _ cli.Command = &InfoCommand{}
}
