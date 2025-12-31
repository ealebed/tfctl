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
	"os"

	"github.com/ealebed/tfctl/pkg/output"
	"github.com/ealebed/tfctl/utils"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// All available service provider types see here:
// https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#ServiceProviderType

// saveOptions represents options for save command
type saveOptions struct {
	*oAuthClientOptions
	providerType  string
	providerToken string
}

// NewOAuthClientSaveCmd returns new OAuth client save command
func NewOAuthClientSaveCmd(oAuthClientOptions *oAuthClientOptions) *cobra.Command {
	options := &saveOptions{
		oAuthClientOptions: oAuthClientOptions,
	}

	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{"create"},
		Short:   "create an OAuth client to connect an organization and a VCS provider",
		Long:    "create an OAuth client to connect an organization and a VCS provider. Expect that only one OAuth client per service provider exist",
		Example: "tfctl OAuthClient save [--providerType=...] [--token=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return saveOAuthClient(cmd, options)
		},
	}

	cmd.Flags().StringVar(&options.providerType, "providerType", "gitlab",
		"terraform OAuth client type for creating. At this time provided support only for 'gitlab' and 'github' providers. Feel free to contribute ))")
	cmd.Flags().StringVarP(&options.providerToken, "token", "t", "",
		"The token string you were given by your VCS provider, e.g. 'ghp_xxxxxxxxxxxxxxx' for a GitHub")
	if err := cmd.MarkFlagRequired("providerType"); err != nil {
		return nil
	}
	if err := cmd.MarkFlagRequired("token"); err != nil {
		return nil
	}

	return cmd
}

func saveOAuthClient(_ *cobra.Command, options *saveOptions) error {
	c := options.TClient
	ctx := context.Background()

	var createOptions tfe.OAuthClientCreateOptions
	switch options.providerType {
	case "github":
		createOptions = tfe.OAuthClientCreateOptions{
			APIURL:          tfe.String("https://api.github.com"),
			HTTPURL:         tfe.String("https://github.com"),
			OAuthToken:      tfe.String(options.providerToken),
			ServiceProvider: tfe.ServiceProvider("github"),
		}
	case "gitlab":
		createOptions = tfe.OAuthClientCreateOptions{
			APIURL:          tfe.String("https://gitlab.com/api/v4"),
			HTTPURL:         tfe.String("https://gitlab.com"),
			OAuthToken:      tfe.String(options.providerToken),
			ServiceProvider: tfe.ServiceProvider("gitlab_hosted"),
		}
	default:
		fmt.Println("Provider type '" + options.providerType +
			"' not supported here. All service provider types see here: " +
			"https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#ServiceProviderType")
		fmt.Println("At this time provided support only for 'gitlab' and 'github' providers.")
		os.Exit(1)
	}

	// List all the OAuth clients for a given organization
	OAuthClients, err := c.OAuthClients.List(ctx, options.TerraformOrganization, &tfe.OAuthClientListOptions{})
	if err != nil {
		return err
	}
	OAuthClientID := utils.GetOAuthClientID(OAuthClients, options.providerType)

	if OAuthClientID == "" {
		// Create an OAuth client to connect an organization and a VCS provider
		OAuthClient, err := c.OAuthClients.Create(ctx, options.TerraformOrganization, createOptions)
		if err != nil {
			return err
		}
		output.JsonPrettyOutput(OAuthClient, "OAuthClient")
	} else {
		fmt.Println("OAuth client for given service provider already exists, check with:\n\t 'tfctl OAuthClient list'\n" +
			"We expect only one client per service provider...")
	}

	return nil
}
