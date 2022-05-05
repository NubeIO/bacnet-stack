package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var Interface string
var Port int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bacnet-stack",
	Short: `A for running bacnet-stack commands`,
	Long:  `A for running bacnet-stack commands`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type flag struct {
	name      string
	shortHand string
	usage     string
}

var flagsCommon = struct {
	iface flag
	port  flag
}{
	iface: flag{name: "interface", shortHand: "f", usage: "Interface e.g. eth0"},
	port:  flag{name: "port", shortHand: "p", usage: "bacnet udp port"},
}

var flagsWhois = struct {
	networkNumber flag
	startRange    flag
	endRange      flag
}{
	networkNumber: flag{name: "network", shortHand: "n", usage: "bacnet network number"},
	startRange:    flag{name: "start", shortHand: "s", usage: "bacnet device id start range"},
	endRange:      flag{name: "end", shortHand: "e", usage: "bacnet device id end range"},
}

var flagsPoint = struct {
	deviceId        flag
	objectType      flag
	objectInstance  flag
	property        flag
	writeValue      flag
	priority        flag
	releasePriority flag
}{
	deviceId:        flag{name: "device", shortHand: "d", usage: "bacnet device number"},
	objectType:      flag{name: "type", shortHand: "t", usage: "bacnet object type ao, av, bv"},
	objectInstance:  flag{name: "inst", shortHand: "i", usage: "bacnet object number 1, 2 as in binary-input-1"},
	property:        flag{name: "prop", shortHand: "r", usage: "what to read example: present value:85"},
	writeValue:      flag{name: "value", shortHand: "v", usage: "what to read example: present value:85"},
	priority:        flag{name: "priority", shortHand: "q", usage: "what to read example: present value:85"},
	releasePriority: flag{name: "release", shortHand: "o", usage: "what to read example: present value:85"},
}

func init() {
	cobra.OnInitialize(initConfig)

	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.baccli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&Interface, flagsCommon.iface.name, flagsCommon.iface.shortHand, "eth0", flagsCommon.iface.usage)
	RootCmd.PersistentFlags().IntVarP(&Port, flagsCommon.port.name, flagsCommon.port.shortHand, int(0xBAC0), flagsCommon.port.usage)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "l", false, "Help message for toggle")

	// We want to allow this to be accessed
	viper.BindPFlag("interface", RootCmd.PersistentFlags().Lookup("interface"))
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".baccli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".baccli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
