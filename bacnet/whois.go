package bacnet

import (
	"bufio"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"strings"
)

/*
Usage: bacwi [device-instance-min [device-instance-max]]
       [--dnet][--dadr][--mac]
       [--version][--help]
Send BACnet WhoIs service request to a device or multiple
devices, and wait for responses. Displays any devices found
and their network information.

device-instance:
BACnet Device Object Instance number that you are trying
to send a Who-Is service request. The value should be in
the range of 0 to 4194303. A range of values can also be
specified by using a minimum value and a maximum value.

--dnet N
BACnet network number N for directed requests.
Valid range is from 0 to 65535 where 0 is the local connection
and 65535 is network broadcast.

--mac A
BACnet mac address.Valid ranges are from 00 to FF (hex) for MS/TP or ARCNET,
or an IP string with optional port number like 10.1.2.3:47808
or an Ethernet MAC in hex like 00:21:70:7e:32:bb

--dadr A
BACnet mac address on the destination BACnet network number.
Valid ranges are from 00 to FF (hex) for MS/TP or ARCNET,
or an IP string with optional port number like 10.1.2.3:47808
or an Ethernet MAC in hex like 00:21:70:7e:32:bb

Send a WhoIs request to DNET 123:
bacwi --dnet 123
Send a WhoIs request to MAC 10.0.0.1 DNET 123 DADR 05h:
bacwi --mac 10.0.0.1 --dnet 123 --dadr 05
Send a WhoIs request to MAC 10.1.2.3:47808:
bacwi --mac 10.1.2.3:47808
Send a WhoIs request to Device 123:
bacwi 123
Send a WhoIs request to Devices from 1000 to 9000:
bacwi 1000 9000
Send a WhoIs request to Devices from 1000 to 9000 on DNET 123:
bacwi 1000 9000 --dnet 123
Send a WhoIs request to all devices:
bacwi
*/

type TypeWhoIs struct {
	Common    *Common
	LowLimit  int
	HighLimit int
}

type ResponseWhoIs struct {
	DeviceID  int
	DeviceMac string
}

func whoIsBuilder(t *TypeWhoIs) (cmd []string) {
	networkNumber := types.ToString(t.Common.NetworkNumber)
	lowLimit := types.ToString(t.LowLimit)
	highLimit := types.ToString(t.HighLimit)
	cmd = []string{whoIs}
	if networkNumber != "-1" {
		cmd = append(cmd, "--dnet", networkNumber)
	}
	if lowLimit != "-1" && highLimit != "-1" {
		cmd = append(cmd, lowLimit, highLimit)
		return
	}
	return

}

func (inst *BACnet) WhoIs(t *TypeWhoIs) {
	path, _ := inst.getBacnetDir()
	out := ""
	cmd := whoIsBuilder(t)

	out, err = inst.Run(path, cmd[0:]...)
	if err != nil {
		//return
	}
	fmt.Println(out)
	scanner := bufio.NewScanner(strings.NewReader(out))
	var bacnetDevice []ResponseWhoIs
	for scanner.Scan() {
		res := scanner.Text()
		res1 := strings.HasPrefix(res, ";")
		if !res1 {
			fmt.Println(res)
			res3 := strings.Fields(res)
			if len(res3) > 3 {
				dev := ResponseWhoIs{}
				dev.DeviceID = types.ToInt(res3[0])
				dev.DeviceMac = res3[1]
				bacnetDevice = append(bacnetDevice, dev)
			}
		}
	}

	for i, dev := range bacnetDevice {
		fmt.Println(i, dev)
	}

}
