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
	"context"

	"github.com/ealebed/tfctl/pkg/output"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// listOptions represents options for list command
type listOptions struct {
	*policySetOptions
}

// NewPolicySetListCmd returns new policy set list command
func NewPolicySetListCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &listOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list all the policy sets for a given organization",
		Long:    "list all the policy sets for a given organization",
		Example: "tfctl policySet list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listPolicySets(cmd, options)
		},
	}

	return cmd
}

func listPolicySets(_ *cobra.Command, options *listOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the policy sets for a given organization
	policySetList, err := c.PolicySets.List(ctx, options.TerraformOrganization, &tfe.PolicySetListOptions{})
	if err != nil {
		return err
	}

	for _, policySet := range policySetList.Items {
		if options.Expand {
			output.JsonOutput(policySet)
		} else {
			output.JsonPrettyOutput(policySet, "policySet")
		}
	}

	return nil
}
