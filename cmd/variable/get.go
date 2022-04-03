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
	"github.com/ealebed/tfctl/utils"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// getOptions represents options for get command
type getOptions struct {
	*variableOptions
	workspaceName string
	variableName  string
}

// NewVariableGetCmd returns new variable get command
func NewVariableGetCmd(variableOptions *variableOptions) *cobra.Command {
	options := &getOptions{
		variableOptions: variableOptions,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"read"},
		Short:   "read a variable by its ID from provided terraform workspace",
		Long:    "read a variable by its ID from provided terraform workspace",
		Example: "tfctl variable get [--workspace=...] [--variable=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getVariable(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform workspace name for getting variable")
	cmd.Flags().StringVarP(&options.variableName, "variable", "v", "", "terraform variable name for getting info")
	cmd.MarkFlagRequired("workspace")
	cmd.MarkFlagRequired("variable")

	return cmd
}

func getVariable(cmd *cobra.Command, options *getOptions) error {
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

	// Read a variable by its ID
	variable, err := c.Variables.Read(ctx, workspace.ID, variableID)
	if err != nil {
		return err
	}

	if options.Expand {
		output.JsonOutput(variable)
	} else {
		output.JsonPrettyOutput(variable, "variable")
	}

	return nil
}
