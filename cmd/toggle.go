/*
Copyright © 2022 JEROEN SMINK <jeroen@sminkware.com>

*/
package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/akominch/yeelight"
	"github.com/spf13/cobra"
)

func newBulb(ip string) *yeelight.Bulb {
	config := yeelight.BulbConfig{
		Ip: ip,
	}
	bulb := yeelight.New(config)

	if bulb == nil {
		fmt.Printf("No Yeelight device found for ip: %s\n", ip)
		os.Exit(1)
	}
	return bulb
}

func checkIpAddress(ip string) {
	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - is not a valid IPv4 Address\n", ip)
		os.Exit(1)
	}
}

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle a Yeelight device based on the IPv4 Address",
	Long:  `Toggle your Yeelight device on and off based ont he given IPv4 Address as parameter flag`,
	Run: func(cmd *cobra.Command, args []string) {

		ip, _ := cmd.Flags().GetString("ip")
		checkIpAddress(ip)

		bulb := newBulb(ip)

		isOn, _ := bulb.IsOn()
		if isOn {
			bulb.TurnOff()
		} else {
			bulb.TurnOn()
		}

	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
	toggleCmd.Flags().String("ip", "", "The IPv4 of the device you want to toggle")
}
