package bacnet

type Common struct {
	NetworkMac     string
	DeviceMac      string //mac mstp-mac
	NetworkNumber  int    //dnet
	DeviceID       int    //bacnet device id
	PrintBacnetRes bool
}

type CommonProps struct {
}
