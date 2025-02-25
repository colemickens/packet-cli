// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"time"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

var (
	projectName     string
	projectID       string
	facility        string
	plan            string
	hostname        string
	operatingSystem string
	billingCycle    string

	storage               string
	userdata              string
	customdata            string
	tags                  []string
	ipxescripturl         string
	publicIPv4SubnetSize  int
	alwaysPXE             bool
	hardwareReservationID string
	spotInstance          bool
	spotPriceMax          float64
	terminationTime       string
)

var createDeviceCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a device",
	Long: `Example:

packet device create --hostname [hostname] --plan [plan] --facility [facility_code] --operating-system [operating_system] --project-id [project_UUID]

`,
	Run: func(cmd *cobra.Command, args []string) {

		request := &packngo.DeviceCreateRequest{
			Hostname:     hostname,
			Plan:         plan,
			Facility:     facility,
			OS:           operatingSystem,
			BillingCycle: billingCycle,
			ProjectID:    projectID,
		}

		if storage != "" {
			request.Storage = storage
		}
		if userdata != "" {
			request.UserData = userdata
		}

		if len(tags) > 0 {
			request.Tags = tags
		}
		if ipxescripturl != "" {
			request.IPXEScriptURL = ipxescripturl
		}
		if publicIPv4SubnetSize != 0 {
			request.PublicIPv4SubnetSize = publicIPv4SubnetSize
		}
		if alwaysPXE {
			request.AlwaysPXE = alwaysPXE
		}

		if hardwareReservationID != "" {
			request.HardwareReservationID = hardwareReservationID
		}

		if spotInstance {
			request.SpotInstance = spotInstance
		}

		if spotPriceMax != 0 {
			request.SpotPriceMax = spotPriceMax
		}

		if terminationTime != "" {
			parsedTime, err := time.Parse(time.RFC3339, terminationTime)
			if err != nil {
				fmt.Printf("Error occured while parsing time string: %s", err.Error())
				return
			}
			request.TerminationTime = &packngo.Timestamp{Time: parsedTime}
		}

		device, _, err := PacknGo.Devices.Create(request)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		header := []string{"ID", "Hostname", "OS", "State", "Created"}
		data := make([][]string, 1)
		data[0] = []string{device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

		output(device, header, &data)

	},
}

func init() {
	createDeviceCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project where the device will be created")
	createDeviceCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the device will be created")
	createDeviceCmd.Flags().StringVarP(&plan, "plan", "P", "", "Name of the plan")
	createDeviceCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "Hostname")
	createDeviceCmd.Flags().StringVarP(&operatingSystem, "operating-system", "o", "", "Operating system name for the device")
	createDeviceCmd.MarkFlagRequired("project-id")
	createDeviceCmd.MarkFlagRequired("facility")
	createDeviceCmd.MarkFlagRequired("plan")
	createDeviceCmd.MarkFlagRequired("hostname")
	createDeviceCmd.MarkFlagRequired("operating-system")

	createDeviceCmd.Flags().StringVarP(&storage, "storage", "s", "", "UUID of the storage")
	createDeviceCmd.Flags().StringVarP(&ipxescripturl, "ipxe-script-url", "i", "", "URL to the iPXE script")
	createDeviceCmd.Flags().StringVarP(&userdata, "userdata", "u", "", "User data")
	createDeviceCmd.Flags().StringVarP(&customdata, "customdata", "c", "", "Custom data")
	createDeviceCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Tags for the device: --tags="tag1,tag2"`)
	createDeviceCmd.Flags().IntVarP(&publicIPv4SubnetSize, "public-ipv4-subnet-size", "v", 0, "Size of the public IPv4 subnet")
	createDeviceCmd.Flags().StringVarP(&hardwareReservationID, "hardware-reservation-id", "r", "", "UUID of the hardware reservation")
	createDeviceCmd.Flags().StringVarP(&billingCycle, "billing-cycle", "b", "hourly", "Billing cycle")
	createDeviceCmd.Flags().BoolVarP(&alwaysPXE, "always-pxe", "a", false, ``)
	createDeviceCmd.Flags().BoolVarP(&spotInstance, "spot-instance", "I", false, `Set the device as a spot instance`)
	createDeviceCmd.Flags().Float64VarP(&spotPriceMax, "spot-price-max", "m", 0, `--spot-price-max=1.2 or -m=1.2`)
	createDeviceCmd.Flags().StringVarP(&terminationTime, "termination-time", "T", "", `Device termination time: --termination-time="15:04:05"`)
	createDeviceCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	createDeviceCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
