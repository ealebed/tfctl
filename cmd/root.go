/*
Copyright © 2022 Yevhen Lebid ealebed@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/ealebed/tfctl/cmd/version"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

type RootOptions struct {
	terraformHostname     string
	TerraformOrganization string
	terraformToken        string
	Expand                bool

	TClient *tfe.Client
}

// NewCmdRoot returns new root command
func NewCmdRoot(outWriter, errWriter io.Writer) (*cobra.Command, *RootOptions) {
	options := &RootOptions{}

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.String(),
	}

	cmd.SetOut(outWriter)
	cmd.SetErr(errWriter)

	// Client flags
	cmd.PersistentFlags().StringVar(&options.terraformHostname, "host", "app.terraform.io", "Terraform Enterprise (Cloud) host")
	cmd.PersistentFlags().StringVar(&options.TerraformOrganization, "org", os.Getenv("TF_ORG"), "Terraform Enterprise (Cloud) organization name")
	cmd.PersistentFlags().StringVar(&options.terraformToken, "token", "", "Terraform Enterprise (Cloud) token")

	// Output flags
	cmd.PersistentFlags().BoolVarP(&options.Expand, "expand", "x", false, "Expand output with all possible values")

	// TODO: add dry-run key for destructive operations
	// cmd.PersistentFlags().BoolVar(&options.DryRun, "dry-run", true, "print output without real changing system configuration")

	// Initialize client
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		config := &tfe.Config{
			Address: fmt.Sprintf("https://%s", options.terraformHostname),
		}

		// If variable 'TF_TOKEN' is empty, try obtain token from credentials file in user ${HOME} directory
		if os.Getenv("TF_TOKEN") != "" {
			config.Token = os.Getenv("TF_TOKEN")
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			credentialsPath := home + "/.terraform.d/credentials.tfrc.json"

			jsonFile, err := os.Open(credentialsPath)
			if err != nil {
				return err
			}
			defer jsonFile.Close()

			var token map[string]interface{}

			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &token)

			cred := token["credentials"].(map[string]interface{})
			tftoken := cred[options.terraformHostname].(map[string]interface{})

			config.Token = tftoken["token"].(string)
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return err
		}

		options.TClient = client

		return nil
	}

	return cmd, options
}
