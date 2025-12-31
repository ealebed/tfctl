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

	"github.com/ealebed/tfctl/pkg/output"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// listOptions represents options for list command
type listOptions struct {
	*variableOptions
	workspaceName string
}

// NewVariableListCmd returns new variable list command
func NewVariableListCmd(variableOptions *variableOptions) *cobra.Command {
	options := &listOptions{
		variableOptions: variableOptions,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list all the variables associated with the given workspace",
		Long:    "list all the variables associated with the given workspace",
		Example: "tfctl variable list [--workspace=..]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listVariables(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform workspace name for getting variables list")
	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}

	return cmd
}

func listVariables(_ *cobra.Command, options *listOptions) error {
	c := options.TClient
	ctx := context.Background()

	// Check if workspace exists and got its ID
	workspace, err := c.Workspaces.Read(ctx, options.TerraformOrganization, options.workspaceName)
	if err != nil {
		return err
	}

	// List all the variables associated with the given workspace
	variableList, err := c.Variables.List(ctx, workspace.ID, &tfe.VariableListOptions{})
	if err != nil {
		return err
	}

	for _, variable := range variableList.Items {
		if options.Expand {
			output.JsonOutput(variable)
		} else {
			output.JsonPrettyOutput(variable, "variable")
		}
	}

	return nil
}
