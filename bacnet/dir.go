package bacnet

import "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/dirs"

func (inst *BACnet) getBacnetDir() (dir string, err error) {
	dir, err = dirs.GetUserHomeDir()
	dir = "/home/aidan/code/bacnet-stack-bacnet-stack-1.0.0/bin"
	return

}
