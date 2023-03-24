// Copyright 2021 Google LLC
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

package certificates

import (
	"fmt"
	"internal/apiclient"

	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/client/certificates"

	"github.com/spf13/cobra"
)

// GetCmd to get integration flow
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certificate details from a region",
	Long:  "Get certificate details from a region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		if id == "" && name == "" {
			return fmt.Errorf("id and name cannot be empty")
		}
		if id != "" && name != "" {
			return fmt.Errorf("id and name both cannot be set")
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if name != "" {
			apiclient.SetPrintOutput(false)
			version, err := certificates.Find(name)
			if err != nil {
				return err
			}
			apiclient.SetPrintOutput(true)
			_, err = certificates.Get(version)
			return err
		} else {
			_, err = certificates.Get(id)
		}
		return

	},
}

var id string

func init() {
	GetCmd.Flags().StringVarP(&id, "id", "i",
		"", "Certificate name (uuid)")
	GetCmd.Flags().StringVarP(&name, "name", "n",
		"", "Certificate display name")
}