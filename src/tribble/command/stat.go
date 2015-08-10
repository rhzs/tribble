package command

import (
	"strings"
	"log"
	"fmt"
	"os/exec"
)

type StatCommand struct {
	Meta
}

func (c *StatCommand) Run(args []string) int {
	// Write your code here
	output, err := exec.Command("vm_stat").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Free Memory: %s", output)

	return 0
}

func (c *StatCommand) Synopsis() string {
	return "Get memory free status"
}

func (c *StatCommand) Help() string {
	helpText := `
Get	free memory
`
	return strings.TrimSpace(helpText)
}
