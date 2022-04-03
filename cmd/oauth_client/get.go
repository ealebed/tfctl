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

package oauth_client

import (
	"context"

	"github.com/ealebed/tfctl/pkg/output"
	"github.com/ealebed/tfctl/utils"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// getOptions represents options for get command
type getOptions struct {
	*oAuthClientOptions
	providerType string
}

// NewOAuthClientGetCmd returns new OAuth client get command
func NewOAuthClientGetCmd(oAuthClientOptions *oAuthClientOptions) *cobra.Command {
	options := &getOptions{
		oAuthClientOptions: oAuthClientOptions,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"read"},
		Short:   "read an OAuth client",
		Long:    "read an OAuth client. Expect that only one OAuth client per service provider exist",
		Example: "tfctl OAuthClient get [--providerType=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getOAuthClient(cmd, options)
		},
	}

	cmd.Flags().StringVar(&options.providerType, "providerType", "gitlab", "terraform OAuth service provider name for getting info")
	cmd.MarkFlagRequired("providerType")

	return cmd
}

func getOAuthClient(cmd *cobra.Command, options *getOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the OAuth clients for a given organization
	OAuthClients, err := c.OAuthClients.List(ctx, options.TerraformOrganization, &tfe.OAuthClientListOptions{})
	if err != nil {
		return err
	}
	OAuthClientID := utils.GetOAuthClientID(OAuthClients, options.providerType)

	// Read an OAuth client by its ID
	OAuthClient, err := c.OAuthClients.Read(ctx, OAuthClientID)
	if err != nil {
		return err
	}

	if options.Expand {
		output.JsonOutput(OAuthClient)
	} else {
		output.JsonPrettyOutput(OAuthClient, "OAuthClient")
	}

	return nil
}
