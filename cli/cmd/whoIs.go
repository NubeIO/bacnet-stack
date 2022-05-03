package cmd

import (
	"github.com/NubeIO/bacnet-stack/bacnet"
	"github.com/spf13/cobra"
)

// Flags
var startRange int
var endRange int
var networkNumber int

// whoIsCmd represents the whoIs command
var whoIsCmd = &cobra.Command{
	Use:   "whois",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: main,
}

func main(cmd *cobra.Command, args []string) {

	//var propInt types.PropertyType
	//// Check to see if an int was passed
	//if i, err := strconv.Atoi(propertyType); err == nil {
	//	propInt = types.PropertyType(uint32(i))
	//} else {
	//	propInt, err = types.Get(propertyType)
	//}
	//
	//if types.IsDeviceProperty(propInt) {
	//	objectType = 8
	//}
	//fmt.Println(propertyType)
	//fmt.Println(types.String(propInt))

	bac := bacnet.New()

	t := &bacnet.TypeWhoIs{
		Common: &bacnet.Common{
			NetworkNumber: networkNumber,
		},
		LowLimit:  startRange,
		HighLimit: endRange,
	}
	bac.WhoIs(t)

}

func init() {
	RootCmd.AddCommand(whoIsCmd)
	whoIsCmd.Flags().IntVarP(&networkNumber, flagsWhois.networkNumber.name, flagsWhois.networkNumber.shortHand, -1, flagsWhois.networkNumber.usage)
	whoIsCmd.Flags().IntVarP(&startRange, flagsWhois.startRange.name, flagsWhois.startRange.shortHand, -1, flagsWhois.startRange.usage)
	whoIsCmd.Flags().IntVarP(&endRange, flagsWhois.endRange.name, flagsWhois.endRange.shortHand, -1, flagsWhois.endRange.usage)

}
