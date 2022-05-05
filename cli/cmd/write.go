package cmd

import (
	"github.com/NubeIO/bacnet-stack/bacnet"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/spf13/cobra"
)

// Flags
//var deviceId int
//var objectType string
//var objectInstance int
//var property int
var writeValue float64
var writePriority int
var releasePriority bool

// writeCmd represents the whoIs command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: write,
}

func write(cmd *cobra.Command, args []string) {

	bac := bacnet.New()

	t := &bacnet.TypeWrite{
		ObjectType:      objectType,
		ObjectInstance:  objectInstance,
		Property:        property,
		WriteValue:      writeValue,
		WritePriority:   writePriority,
		ReleasePriority: nils.NewBool(releasePriority),

		Common: &bacnet.Common{
			DeviceID:      deviceId,
			NetworkNumber: networkNumber,
		},
	}
	bac.Write(t)

}

func init() {
	//write example
	//go run main.go write --device=3000 --type=4 --inst=1 --priority=16 --value=1 --release=false
	RootCmd.AddCommand(writeCmd)
	writeCmd.Flags().IntVarP(&deviceId, flagsPoint.deviceId.name, flagsPoint.deviceId.shortHand, 2508, flagsPoint.deviceId.usage)
	writeCmd.Flags().StringVarP(&objectType, flagsPoint.objectType.name, flagsPoint.objectType.shortHand, "", flagsPoint.objectType.usage)
	writeCmd.Flags().IntVarP(&objectInstance, flagsPoint.objectInstance.name, flagsPoint.objectInstance.shortHand, 1, flagsPoint.objectInstance.usage)
	writeCmd.Flags().IntVarP(&property, flagsPoint.property.name, flagsPoint.property.shortHand, 1, flagsPoint.property.usage)
	writeCmd.Flags().Float64VarP(&writeValue, flagsPoint.writeValue.name, flagsPoint.writeValue.shortHand, 1, flagsPoint.writeValue.usage)
	writeCmd.Flags().IntVarP(&writePriority, flagsPoint.priority.name, flagsPoint.priority.shortHand, 1, flagsPoint.priority.usage)
	writeCmd.Flags().BoolVarP(&releasePriority, flagsPoint.releasePriority.name, flagsPoint.releasePriority.shortHand, false, flagsPoint.releasePriority.usage)

}
