package command

import (
	"strings"
	"fmt"

	"github.com/klauspost/cpuid"
)

type InfoCommand struct {
	Meta
}

func (c *InfoCommand) Run(args []string) int {
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model)
	fmt.Println("Features:", cpuid.CPU.Features)
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")

	// Test if we have a specific feature:
	if cpuid.CPU.SSE() {
		fmt.Println("We have Streaming SIMD Extensions")
	}
	return 0
}

func (c *InfoCommand) Synopsis() string {
	return "Get cpu information"
}

func (c *InfoCommand) Help() string {
	helpText := `
Get	cpu information
`
	return strings.TrimSpace(helpText)
}

