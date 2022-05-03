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
	//go run main.go read -d=3000 -t=1 -i=1 -r=85
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&deviceId, flagsRead.deviceId.name, flagsRead.deviceId.shortHand, 2508, flagsRead.deviceId.usage)
	readCmd.Flags().StringVarP(&objectType, flagsRead.objectType.name, flagsRead.objectType.shortHand, "", flagsRead.objectType.usage)
	readCmd.Flags().IntVarP(&objectInstance, flagsRead.objectInstance.name, flagsRead.objectInstance.shortHand, 1, flagsRead.objectInstance.usage)
	readCmd.Flags().IntVarP(&property, flagsRead.property.name, flagsRead.property.shortHand, 1, flagsRead.property.usage)

}
