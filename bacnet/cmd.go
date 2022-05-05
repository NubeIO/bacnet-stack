package bacnet

import (
	"errors"
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

//bacnet stack commands
const (
	whoIs = "./bacwi"
	read  = "./bacrp"
	write = "./bacwp"
)

//bacnet stack error messages
const (
	apduTimeout              = "APDU Timeout!"
	errorUnknownProp         = "unknown-property"
	deviceOperationalProblem = "device: operational-problem"
	WwAcknowledged           = "WriteProperty Acknowledged!"
)

func resError(in string) error {
	if strings.Contains(in, apduTimeout) {
		return errors.New("point or device timeout")
	} else if strings.Contains(in, errorUnknownProp) {
		return errors.New("device unknown-property")
	} else if strings.Contains(in, deviceOperationalProblem) {
		return errors.New("device operational-problem")
	}

	return nil
}

// Run runs given command with parameters and return combined output
func (inst *BACnet) Run(path string, cmdAndParams ...string) (string, error) {
	if len(cmdAndParams) <= 0 {
		return "", fmt.Errorf("no command provided")
	}
	fmt.Println(exec.Command(cmdAndParams[0], cmdAndParams[1:]...).String())
	cmd := exec.Command(cmdAndParams[0], cmdAndParams[1:]...)
	cmd.Dir = path
	output, err := cmd.Output()
	outAsString := strings.TrimRight(string(output), "\n")
	err = resError(string(output))
	if err != nil {
		fmt.Println("BACNet cmd err", err)
		return "", err
	} else {
		fmt.Println("BACNet cmd res", outAsString)
	}
	return outAsString, err
}
