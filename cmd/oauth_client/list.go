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
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// listOptions represents options for list command
type listOptions struct {
	*oAuthClientOptions
}

// NewOAuthClientListCmd returns new OAuth client list command
func NewOAuthClientListCmd(oAuthClientOptions *oAuthClientOptions) *cobra.Command {
	options := &listOptions{
		oAuthClientOptions: oAuthClientOptions,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list all the OAuth clients for a given organization",
		Long:    "list all the OAuth clients for a given organization",
		Example: "tfctl OAuthClient list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listOAuthClients(cmd, options)
		},
	}

	return cmd
}

func listOAuthClients(cmd *cobra.Command, options *listOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the OAuth clients for a given organization
	OAuthClients, err := c.OAuthClients.List(ctx, options.TerraformOrganization, &tfe.OAuthClientListOptions{})
	if err != nil {
		return err
	}

	for _, OAuthClient := range OAuthClients.Items {
		if options.Expand {
			output.JsonOutput(OAuthClient)
		} else {
			output.JsonPrettyOutput(OAuthClient, "OAuthClient")
		}
	}

	return nil
}
