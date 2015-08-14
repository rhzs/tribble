package command

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	pid int
	cpu float64
}

type CpuCommand struct {
	Meta
}

func (c *CpuCommand) Run(args []string) int {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err!=nil {
			break;
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range(tokens) {
			if t!="" && t!="\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err!=nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err!=nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range(processes) {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
	return 0
}

func (c *CpuCommand) Synopsis() string {
	return "Get cpu usage"
}

func (c *CpuCommand) Help() string {
	helpText := `
Get	cpu usage
`
	return strings.TrimSpace(helpText)
}

