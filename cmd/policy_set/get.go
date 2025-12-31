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
	"github.com/ealebed/tfctl/utils"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// getOptions represents options for get command
type getOptions struct {
	*policySetOptions
	policySetName string
}

// NewPolicySetGetCmd returns new policy set get command
func NewPolicySetGetCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &getOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"read"},
		Short:   "read a policy set by its ID from provided terraform organization",
		Long:    "read a policy set by its ID from provided terraform organization",
		Example: "tfctl policySet get [--policySet=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getPolicySet(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.policySetName, "policySet", "p", "", "terraform policy set name for getting info")
	if err := cmd.MarkFlagRequired("policySet"); err != nil {
		return nil
	}

	return cmd
}

func getPolicySet(_ *cobra.Command, options *getOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the policy sets for a given organization and filter policy set ID by provided name
	policySetList, err := c.PolicySets.List(ctx, options.TerraformOrganization, &tfe.PolicySetListOptions{})
	if err != nil {
		return err
	}
	policySetID := utils.GetPolicySetID(policySetList, options.policySetName)

	// Read a policy set by its ID
	policySet, err := c.PolicySets.Read(ctx, policySetID)
	if err != nil {
		return err
	}

	if options.Expand {
		output.JsonOutput(policySet)
	} else {
		output.JsonPrettyOutput(policySet, "policySet")
	}

	return nil
}
