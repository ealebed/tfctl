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

//nolint:dupl // This file has similar command setup pattern to save.go
package workspace

import (
	"context"

	"github.com/ealebed/tfctl/pkg/output"

	"github.com/spf13/cobra"
)

// getOptions represents options for get command
type getOptions struct {
	*workspaceOptions
	workspaceName string
}

// NewWorkspaceGetCmd returns new workspaces list command
func NewWorkspaceGetCmd(workspaceOptions *workspaceOptions) *cobra.Command {
	options := &getOptions{
		workspaceOptions: workspaceOptions,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"read"},
		Short:   "read a workspace by its name and organization name",
		Long:    "read a workspace by its name and organization name",
		Example: "tfctl ws get [--workspace=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getWorkspace(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform workspace name for getting info")
	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}

	return cmd
}

func getWorkspace(_ *cobra.Command, options *getOptions) error {
	c := options.TClient
	ctx := context.Background()

	// Read a workspace by its name and organization name
	workspace, err := c.Workspaces.Read(ctx, options.TerraformOrganization, options.workspaceName)
	if err != nil {
		return err
	}

	if options.Expand {
		output.JsonOutput(workspace)
	} else {
		output.JsonPrettyOutput(workspace, "workspace")
	}

	return nil
}
