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

package workspace

import (
	"github.com/ealebed/tfctl/cmd"

	"github.com/spf13/cobra"
)

// Package describes the workspace related methods that the Terraform
// Enterprise API supports.
//
// TFE API docs: https://www.terraform.io/docs/cloud/api/workspaces.html

type workspaceOptions struct {
	*cmd.RootOptions
}

// NewWorkspaceCmd create new workspace command
func NewWorkspaceCmd(rootOptions *cmd.RootOptions) *cobra.Command {
	options := &workspaceOptions{
		RootOptions: rootOptions,
	}

	cobraCmd := &cobra.Command{
		Use:     "ws",
		Aliases: []string{"workspace"},
		Short:   "Work with terraform workspaces",
		Long:    "Work with terraform workspaces",
		Example: "",
	}

	// create subcommands
	cobraCmd.AddCommand(NewWorkspaceGetCmd(options))
	cobraCmd.AddCommand(NewWorkspaceListCmd(options))
	cobraCmd.AddCommand(NewWorkspaceSaveCmd(options))
	cobraCmd.AddCommand(NewWorkspaceDeleteCmd(options))

	return cobraCmd
}
