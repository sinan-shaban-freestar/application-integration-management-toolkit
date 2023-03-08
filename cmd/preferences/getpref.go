// Copyright 2020 Google LLC
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

package preferences

import (
	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/apiclient"
	"github.com/spf13/cobra"
)

// Cmd to get org details
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get preferences for integrationcli",
	Long:  "Get preferences for integrationcli",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if err = apiclient.GetPreferences(); err != nil {
			return err
		}

		return nil
	},
}
