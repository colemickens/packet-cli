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

	"github.com/spf13/cobra"
)

// plansCmd represents the plans command
var retrievePlansCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a list of all available plans.",
	Long: `Example:

  packet plans get
  
  `,
	Run: func(cmd *cobra.Command, args []string) {
		plans, _, err := PacknGo.Plans.List()
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, len(plans))

		for i, p := range plans {
			data[i] = []string{p.ID, p.Slug, p.Name}
		}
		header := []string{"ID", "Slug", "Name"}

		output(plans, header, &data)
	},
}

func init() {
	retrievePlansCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrievePlansCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
