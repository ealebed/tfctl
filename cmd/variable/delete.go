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

package variable

import (
	"context"
	"fmt"

	"github.com/ealebed/tfctl/utils"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// deleteOptions represents options for delete command
type deleteOptions struct {
	*variableOptions
	workspaceName string
	variableName  string
}

// NewVariableDeleteCmd returns new variable delete command
func NewVariableDeleteCmd(variableOptions *variableOptions) *cobra.Command {
	options := &deleteOptions{
		variableOptions: variableOptions,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "delete variable from provided terraform workspace",
		Long:    "delete variable from provided terraform workspace",
		Example: "tfctl variable delete [--workspace=...] [--variable=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteVariable(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform workspace name for deleting variable")
	cmd.Flags().StringVarP(&options.variableName, "variable", "v", "", "terraform variable name to delete")
	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}
	if err := cmd.MarkFlagRequired("variable"); err != nil {
		return nil
	}

	return cmd
}

func deleteVariable(_ *cobra.Command, options *deleteOptions) error {
	c := options.TClient
	ctx := context.Background()

	// Check if workspace exists and got its ID
	workspace, err := c.Workspaces.Read(ctx, options.TerraformOrganization, options.workspaceName)
	if err != nil {
		return err
	}

	// List all the variables associated with the given workspace and filter variable's ID by provided name
	variableList, err := c.Variables.List(ctx, workspace.ID, &tfe.VariableListOptions{})
	if err != nil {
		return err
	}
	variableID := utils.GetVariableID(variableList, options.variableName)

	// Delete a variable by its ID
	if err := c.Variables.Delete(ctx, workspace.ID, variableID); err != nil {
		return err
	}
	fmt.Println("Variable '" + options.variableName + "' from workspace '" + options.workspaceName + "' deleted successfully!")

	return nil
}
