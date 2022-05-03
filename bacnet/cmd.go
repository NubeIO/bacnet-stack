package bacnet

import (
	"fmt"
	"os/exec"
	"strings"
)

type BACnet struct {
}

func New() *BACnet {
	return &BACnet{}
}

var err error

// Run runs given command with parameters and return combined output
func (inst *BACnet) Run(path string, cmdAndParams ...string) (string, error) {
	if len(cmdAndParams) <= 0 {
		return "", fmt.Errorf("no command provided")
	}
	fmt.Println(exec.Command(cmdAndParams[0], cmdAndParams[1:]...).String())
	//output, err := exec.Command(cmdAndParams[0], cmdAndParams[1:]...).CombinedOutput()
	cmd := exec.Command(cmdAndParams[0], cmdAndParams[1:]...)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(output), "\n"), err
}

const (
	whoIs = "./bacwi"
	read  = "./bacrp"
)
