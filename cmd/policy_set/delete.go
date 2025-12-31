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
	"fmt"

	"github.com/ealebed/tfctl/utils"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// deleteOptions represents options for delete command
type deleteOptions struct {
	*policySetOptions
	policySetName string
}

// NewPolicySetDeleteCmd returns new policy set delete command
func NewPolicySetDeleteCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &deleteOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "delete a policy set from provided organization",
		Long:    "delete a policy set from provided organization",
		Example: "tfctl policySet delete [--policySet=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deletePolicySet(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.policySetName, "policySet", "p", "", "terraform policy set ID for deleting")
	if err := cmd.MarkFlagRequired("policySet"); err != nil {
		return nil
	}

	return cmd
}

func deletePolicySet(_ *cobra.Command, options *deleteOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the policy sets for a given organization and filter policy set ID by provided name
	policySetList, err := c.PolicySets.List(ctx, options.TerraformOrganization, &tfe.PolicySetListOptions{})
	if err != nil {
		return err
	}
	policySetID := utils.GetPolicySetID(policySetList, options.policySetName)

	// Delete a policy set by its ID
	if err := c.PolicySets.Delete(ctx, policySetID); err != nil {
		return err
	}
	fmt.Println("Policy set '" + options.policySetName + "' deleted successfully!")

	return nil
}
