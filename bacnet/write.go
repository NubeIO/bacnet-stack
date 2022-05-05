package bacnet

import (
	"errors"
	"fmt"
	btypes "github.com/NubeIO/bacnet-stack/types"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
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

type TypeWrite struct {
	Common          *Common
	ObjectType      string
	ObjectInstance  int
	Property        int
	WriteValue      float64
	WritePriority   int
	ReleasePriority *bool
}

type ResponseWrite struct {
	DeviceID  int
	DeviceMac string
}

const (
	typeBool  = "9"
	typeFloat = "4"
	typeNull  = "0"
)

func writeBuilder(t *TypeWrite) (cmd []string) {
	deviceId := t.Common.DeviceID
	cmd = []string{write}
	value, dataType, priority, err := writeBuilderData(t)
	if err != nil {
		fmt.Println("error on bacnet write parseWrite err", err)
	}
	//change to null
	if nils.BoolIsNil(t.ReleasePriority) {
		dataType = "0"
		value = 0
	}
	//./bacwp 123 1 1 85 16 -1 4 11.11
	// ./bacwp 123 1 1 85 16 -1 0 0 write null
	cmd = append(cmd, types.ToString(deviceId), t.ObjectType, types.ToString(t.ObjectInstance), "85", fmt.Sprintf("%d", priority), "-1", dataType, fmt.Sprintf("%f", value))
	return
}

func writeBuilderData(t *TypeWrite) (out float64, dataType string, priority int, err error) {
	//Binary-Output-4 Binary-Value-5 Analog-Output-1 Analog-Value-2
	value := t.WriteValue
	priority = t.WritePriority
	typ := t.ObjectType
	if priority < 1 {
		fmt.Println("priority must be between 1 and 16")
		priority = 1
	} else if priority > 16 {
		fmt.Println("priority must be between 1 and 16")
		priority = 16
	}
	if typ == fmt.Sprintf("%d", btypes.BinaryOutputNum) || typ == fmt.Sprintf("%d", btypes.BinaryValueNum) {
		if value > 0 {
			return 1, typeBool, priority, nil
		}
		return 0, typeBool, priority, nil
	} else if typ == fmt.Sprintf("%d", btypes.AnalogOutputNum) || typ == fmt.Sprintf("%d", btypes.AnalogValueNum) {
		return value, typeFloat, priority, nil
	} else {
		return 0, "", priority, errors.New("unknown type")
	}
}

func writeOk(msg string) (ok bool) {
	if msg == WwAcknowledged {
		ok = true
	}
	return
}

/*
write a bo
./bacwp 123 4 1 85 16 -1 9 0

write a bo value to @16 to null
./bacwp 2508 4 1 85 16 -1 0 0
*/

//Write
func (inst *BACnet) Write(t *TypeWrite) (ok bool, err error) {
	path, _ := inst.getBacnetDir()
	out := ""
	cmd := writeBuilder(t)
	out, err = inst.Run(path, cmd[0:]...)
	if err != nil {
		return false, err
	} else {
		ok = writeOk(out)
		return
	}

}
