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
	"github.com/ealebed/tfctl/cmd"

	"github.com/spf13/cobra"
)

// Package describes the variable related methods that the Terraform
// Enterprise API supports.
//
// TFE API docs: https://www.terraform.io/docs/cloud/api/workspace-variables.html

type variableOptions struct {
	*cmd.RootOptions
}

// NewVariableCmd create new variable command
func NewVariableCmd(rootOptions *cmd.RootOptions) *cobra.Command {
	options := &variableOptions{
		RootOptions: rootOptions,
	}

	cobraCmd := &cobra.Command{
		Use:     "variable",
		Aliases: []string{"var"},
		Short:   "Work with terraform variables",
		Long:    "Work with terraform variables",
		Example: "",
	}

	// create subcommands
	cobraCmd.AddCommand(NewVariableListCmd(options))
	cobraCmd.AddCommand(NewVariableGetCmd(options))
	cobraCmd.AddCommand(NewVariableSaveCmd(options))
	cobraCmd.AddCommand(NewVariableDeleteCmd(options))

	return cobraCmd
}
