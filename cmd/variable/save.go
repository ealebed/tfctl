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

// saveOptions represents options for save command
type saveOptions struct {
	*variableOptions
	workspaceName string
	key           string
	value         string
	category      string
	hcl           bool
	sensitive     bool
}

// NewVariableSaveCmd returns new variable save command
func NewVariableSaveCmd(variableOptions *variableOptions) *cobra.Command {
	options := &saveOptions{
		variableOptions: variableOptions,
	}

	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{"create"},
		Short:   "save (create) variable in provided terraform workspace",
		Long:    "save (create or update if already exists) variable in provided terraform workspace",
		Example: "tfctl variable save [--workspace=...] [--key=...] [--value=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return saveVariable(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform workspace name for creating variable")
	cmd.Flags().StringVar(&options.key, "key", "", "terraform variable name to create")
	cmd.Flags().StringVar(&options.value, "value", "", "terraform variable value to create")
	cmd.Flags().StringVar(&options.category, "category", "terraform", "Optional: represents a category type (terraform or env)")
	cmd.Flags().BoolVar(&options.hcl, "hcl", false, "Optional: Whether to evaluate the value of the variable as a string of HCL code")
	cmd.Flags().BoolVar(&options.sensitive, "sensitive", false, "Optional: Whether the value is sensitive")

	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}
	if err := cmd.MarkFlagRequired("key"); err != nil {
		return nil
	}
	if err := cmd.MarkFlagRequired("value"); err != nil {
		return nil
	}

	return cmd
}

func saveVariable(_ *cobra.Command, options *saveOptions) error {
	c := options.TClient
	ctx := context.Background()

	var variable *tfe.Variable

	// Fill needed create options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#VariableCreateOptions
	variableCreateOpts := tfe.VariableCreateOptions{
		Key:       tfe.String(options.key),
		Value:     tfe.String(options.value),
		Category:  tfe.Category(tfe.CategoryType(options.category)),
		HCL:       tfe.Bool(options.hcl),
		Sensitive: tfe.Bool(options.sensitive),
	}

	// Fill needed update options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#VariableUpdateOptions
	variableUpdateOptions := tfe.VariableUpdateOptions{
		Key:       tfe.String(options.key),
		Value:     tfe.String(options.value),
		HCL:       tfe.Bool(options.hcl),
		Sensitive: tfe.Bool(options.sensitive),
	}

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
	variableID := utils.GetVariableID(variableList, options.key)

	if variableID == "" {
		// Create is used to create a new variable
		variable, err = c.Variables.Create(ctx, workspace.ID, variableCreateOpts)
		if err != nil {
			return err
		}
	} else {
		// Update values of an existing variable
		variable, err = c.Variables.Update(ctx, workspace.ID, variableID, variableUpdateOptions)
		if err != nil {
			return err
		}
	}

	output.JsonPrettyOutput(variable, "variable")

	return nil
}
