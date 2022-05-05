package cmd

import (
	"github.com/NubeIO/bacnet-stack/bacnet"
	"github.com/spf13/cobra"
)

// Flags
var deviceId int
var objectType string
var objectInstance int
var property int

// whoIsCmd represents the whoIs command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: read,
}

func read(cmd *cobra.Command, args []string) {
	bac := bacnet.New()
	t := &bacnet.TypeRead{
		ObjectType:     objectType,
		ObjectInstance: objectInstance,
		Property:       property,
		Common: &bacnet.Common{
			DeviceID:      deviceId,
			NetworkNumber: networkNumber,
		},
	}
	bac.Read(t)

}

func init() {
	//read example
	//go run main.go read --device=3000 --type=4 --inst=1 --prop=85
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&deviceId, flagsPoint.deviceId.name, flagsPoint.deviceId.shortHand, 2508, flagsPoint.deviceId.usage)
	readCmd.Flags().StringVarP(&objectType, flagsPoint.objectType.name, flagsPoint.objectType.shortHand, "", flagsPoint.objectType.usage)
	readCmd.Flags().IntVarP(&objectInstance, flagsPoint.objectInstance.name, flagsPoint.objectInstance.shortHand, 1, flagsPoint.objectInstance.usage)
	readCmd.Flags().IntVarP(&property, flagsPoint.property.name, flagsPoint.property.shortHand, 1, flagsPoint.property.usage)

}
