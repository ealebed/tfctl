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
	"fmt"

	"github.com/ealebed/tfctl/utils"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// deleteOptions represents options for delete command
type deleteOptions struct {
	*oAuthClientOptions
	providerType string
}

// NewOAuthClientDeleteCmd returns new OAuth client delete command
func NewOAuthClientDeleteCmd(oAuthClientOptions *oAuthClientOptions) *cobra.Command {
	options := &deleteOptions{
		oAuthClientOptions: oAuthClientOptions,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "delete an OAuth client",
		Long:    "delete an OAuth client. Expect that only one OAuth client per service provider exist",
		Example: "tfctl OAuthClient delete [--providerType=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteOAuthClient(cmd, options)
		},
	}

	cmd.Flags().StringVar(&options.providerType, "providerType", "gitlab", "terraform OAuth service provider name for deleting")
	cmd.MarkFlagRequired("providerType")

	return cmd
}

func deleteOAuthClient(cmd *cobra.Command, options *deleteOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the OAuth clients for a given organization
	OAuthClients, err := c.OAuthClients.List(ctx, options.TerraformOrganization, &tfe.OAuthClientListOptions{})
	if err != nil {
		return err
	}
	OAuthClientID := utils.GetOAuthClientID(OAuthClients, options.providerType)

	// Delete an OAuth client by its ID
	if err := c.OAuthClients.Delete(ctx, OAuthClientID); err != nil {
		return err
	} else {
		fmt.Println("OAuth Client '" + options.providerType + "' deleted successfully!")
	}

	return nil
}
