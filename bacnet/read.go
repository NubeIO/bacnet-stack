package bacnet

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
)

/*
Usage: bacrp device-instance object-type object-instance property [index]
       [--dnet][--dadr][--mac]
       [--version][--help]
Read a property from an object in a BACnet device
and print the value.
--mac A
Optional BACnet mac address.Valid ranges are from 00 to FF (hex) for MS/TP or ARCNET,
or an IP string with optional port number like 10.1.2.3:47808
or an Ethernet MAC in hex like 00:21:70:7e:32:bb

--dnet N
Optional BACnet network number N for directed requests.
Valid range is from 0 to 65535 where 0 is the local connection
and 65535 is network broadcast.

--dadr A
Optional BACnet mac address on the destination BACnet network number.
Valid ranges are from 00 to FF (hex) for MS/TP or ARCNET,
or an IP string with optional port number like 10.1.2.3:47808
or an Ethernet MAC in hex like 00:21:70:7e:32:bb

device-instance:
BACnet Device Object Instance number that you are
trying to communicate to.  This number will be used
to try and bind with the device using Who-Is and
I-Am services.  For example, if you were reading
Device Object 123, the device-instance would be 123.

object-type:
The object type is object that you are reading. It
can be defined either as the object-type name string
as defined in the BACnet specification, or as the
integer value of the enumeration BACNET_OBJECT_TYPE
in bacenum.h. For example if you were reading Analog
Output 2, the object-type would be analog-output or 1.

object-instance:
This is the object instance number of the object that
you are reading.  For example, if you were reading
Analog Output 2, the object-instance would be 2.

property:
The property of the object that you are reading. It
can be defined either as the property name string as
defined in the BACnet specification, or as an integer
value of the enumeration BACNET_PROPERTY_ID in
bacenum.h. For example, if you were reading the Present
Value property, use present-value or 85 as the property.

index:
This integer parameter is the index number of an array.
If the property is an array, individual elements can
be read.  If this parameter is missing and the property
is an array, the entire array will be read.

Example:
If you want read the Present-Value of Analog Output 101
in Device 123, you could send either of the following
commands:
bacrp 123 analog-output 101 present-value
bacrp 123 1 101 85
If you want read the Priority-Array of Analog Output 101
in Device 123, you could send either of the following
commands:
bacrp 123 analog-output 101 priority-array
bacrp 123 1 101 87
*/

type TypeRead struct {
	Common         *Common
	ObjectType     string
	ObjectInstance int
	Property       int
}

type ResponseRead struct {
	DeviceID  int
	DeviceMac string
}

func readBuilder(t *TypeRead) (cmd []string) {
	deviceId := t.Common.DeviceID

	cmd = []string{read}
	//bacrp 123 analog-output 101 priority-array
	cmd = append(cmd, types.ToString(deviceId), t.ObjectType, types.ToString(t.ObjectInstance), types.ToString(t.Property))
	return

}

func (inst *BACnet) Read(t *TypeRead) {
	path, _ := inst.getBacnetDir()
	out := ""
	cmd := readBuilder(t)
	out, err = inst.Run(path, cmd[0:]...)
	if err != nil {
		//return
	}

	aa := types.ToFloat64(out)
	fmt.Println(aa)

}
