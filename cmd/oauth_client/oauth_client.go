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
	"github.com/ealebed/tfctl/cmd"
	"github.com/spf13/cobra"
)

// Package describes the OAuth client related methods that the
// Terraform Enterprise API supports.
//
// TFE API docs:
// https://www.terraform.io/docs/enterprise/api/oauth-clients.html

type oAuthClientOptions struct {
	*cmd.RootOptions
}

// NewOAuthClientCmd create new OAuth client command
func NewOAuthClientCmd(rootOptions *cmd.RootOptions) *cobra.Command {
	options := &oAuthClientOptions{
		RootOptions: rootOptions,
	}

	cmd := &cobra.Command{
		Use:     "OAuthClient",
		Short:   "Work with terraform OAuth clients",
		Long:    "Work with terraform OAuth clients",
		Example: "",
	}

	// create subcommands
	cmd.AddCommand(NewOAuthClientListCmd(options))
	cmd.AddCommand(NewOAuthClientGetCmd(options))
	cmd.AddCommand(NewOAuthClientSaveCmd(options))
	cmd.AddCommand(NewOAuthClientDeleteCmd(options))

	return cmd
}
