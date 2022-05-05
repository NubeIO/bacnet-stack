package bacnet

/*
Analog Input	0
Analog Output	1
Analog Value	2
Binary Input	3
Binary Output	4
Binary Value	5
Calendar	6
Command	7
Device	8
Event Enrollment	9
File	10
Group	11
Loop	12
Multistate Input	13
Multistate Output	14
Notification Class	15
Program	16
Schedule	17
*/

/*
objectIdentifier	75
objectName	77
objectType	79
systemStatus	112
vendorName	121
vendorIdentifier	120
modelName	70
firmwareRevision	44
applicationSoftwareVersion	12
location
description
protocolVersion	98
ProtocolConformanceClass	95
protocolServicesSupported	97
protocolObjectTypesSupported	96
objectList	76
*/

/*
object list
./bacrp 123 8 123 76

device name
./bacrp 123 8 123 28

*/

type Common struct {
	NetworkMac     string
	DeviceMac      string //mac mstp-mac
	NetworkNumber  int    //dnet
	DeviceID       int    //bacnet device id
	PrintBacnetRes bool
}

type CommonProps struct {
}
