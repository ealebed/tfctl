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
	"context"

	"github.com/ealebed/tfctl/pkg/output"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// listOptions represents options for list command
type listOptions struct {
	*workspaceOptions
}

// NewWorkspaceListCmd returns new workspaces list command
func NewWorkspaceListCmd(workspaceOptions *workspaceOptions) *cobra.Command {
	options := &listOptions{
		workspaceOptions: workspaceOptions,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list all the workspaces within an organization",
		Long:    "list all the workspaces within an organization",
		Example: "tfctl ws list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listWorkspaces(cmd, options)
		},
	}

	return cmd
}

func listWorkspaces(_ *cobra.Command, options *listOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the workspaces within an organization
	workspaceList, err := c.Workspaces.List(ctx, options.TerraformOrganization, &tfe.WorkspaceListOptions{})
	if err != nil {
		return err
	}

	for _, workspace := range workspaceList.Items {
		if options.Expand {
			output.JsonOutput(workspace)
		} else {
			output.JsonPrettyOutput(workspace, "workspace")
		}
	}

	return nil
}
