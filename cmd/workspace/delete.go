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
	"fmt"

	"github.com/spf13/cobra"
)

// deleteOptions represents options for delete command
type deleteOptions struct {
	*workspaceOptions
	workspaceName string
}

// NewWorkspaceDeleteCmd returns new workspace delete command
func NewWorkspaceDeleteCmd(workspaceOptions *workspaceOptions) *cobra.Command {
	options := &deleteOptions{
		workspaceOptions: workspaceOptions,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "delete a terraform workspace by its name",
		Long:    "delete a terraform workspace by its name",
		Example: "tfctl ws delete [--workspace=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteWorkspace(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "name terraform workspace to delete")
	cmd.MarkFlagRequired("workspace")

	return cmd
}

func deleteWorkspace(cmd *cobra.Command, options *deleteOptions) error {
	c := options.TClient
	ctx := context.Background()

	// Delete a workspace by its name
	if err := c.Workspaces.Delete(ctx, options.TerraformOrganization, options.workspaceName); err != nil {
		return err
	} else {
		fmt.Println("Workspace '" + options.workspaceName + "' deleted successfully!")
	}

	return nil
}
