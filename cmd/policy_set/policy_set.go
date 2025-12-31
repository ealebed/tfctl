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

package policy_set

import (
	"github.com/ealebed/tfctl/cmd"

	"github.com/spf13/cobra"
)

// Package describes the policy set related methods that the Terraform
// Enterprise API supports.
//
// TFE API docs: https://www.terraform.io/docs/cloud/api/policy-sets.html

type policySetOptions struct {
	*cmd.RootOptions
}

// NewPolicySetCmd create new policy set command
func NewPolicySetCmd(rootOptions *cmd.RootOptions) *cobra.Command {
	options := &policySetOptions{
		RootOptions: rootOptions,
	}

	cobraCmd := &cobra.Command{
		Use:     "policySet",
		Aliases: []string{"ps"},
		Short:   "Work with terraform policy sets",
		Long:    "Work with terraform policy sets (managing individual policies is deprecated)",
		Example: "",
	}

	// create subcommands
	cobraCmd.AddCommand(NewPolicySetListCmd(options))
	cobraCmd.AddCommand(NewPolicySetGetCmd(options))
	cobraCmd.AddCommand(NewPolicySetSaveCmd(options))
	cobraCmd.AddCommand(NewPolicySetDeleteCmd(options))
	cobraCmd.AddCommand(NewPolicySetAttachCmd(options))
	cobraCmd.AddCommand(NewPolicySetDetachCmd(options))

	return cobraCmd
}
