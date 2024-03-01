// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectors

import (
	"internal/apiclient"
	"internal/client/connections"

	"github.com/spf13/cobra"
)

// DelCustomCmd to get connection
var DelCustomCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a custom connection",
	Long:  "Delete a custom connection",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		cmdProject := cmd.Flag("proj")
		cmdRegion := cmd.Flag("reg")

		if err = apiclient.SetRegion(cmdRegion.Value.String()); err != nil {
			return err
		}
		return apiclient.SetProjectID(cmdProject.Value.String())
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		name := cmd.Flag("name").Value.String()
		_, err = connections.DeleteCustom(name, force)
		return err
	},
}

var force bool

func init() {
	var name string

	DelCustomCmd.Flags().StringVarP(&name, "name", "n",
		"", "The name of the custom connection")

	DelCustomCmd.Flags().BoolVarP(&force, "force", "",
		false, "Force delete the custom connection")

	_ = DelCustomCmd.MarkFlagRequired("name")
}
