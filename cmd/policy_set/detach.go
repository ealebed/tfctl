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

// detachOptions represents options for detach command
type detachOptions struct {
	*policySetOptions
	policySetName string
	workspaceName string
}

// NewPolicySetDetachCmd returns new policy set detach command
func NewPolicySetDetachCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &detachOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "detach",
		Short:   "remove workspace(s) from a policy set in terraform organization",
		Long:    "remove workspace(s) from a policy set in terraform organization",
		Example: "tfctl policySet detach [--policySet=...] [--workspace=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return detachFromPolicySet(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.policySetName, "policySet", "p", "", "terraform policy set name")
	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform policy workspace name for detaching from a policy set")
	if err := cmd.MarkFlagRequired("policySet"); err != nil {
		return nil
	}
	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}

	return cmd
}

func detachFromPolicySet(_ *cobra.Command, options *detachOptions) error {
	c := options.TClient
	ctx := context.Background()

	// List all the policy sets for a given organization and filter policy set ID by provided name
	policySetList, err := c.PolicySets.List(ctx, options.TerraformOrganization, &tfe.PolicySetListOptions{})
	if err != nil {
		return err
	}
	policySetID := utils.GetPolicySetID(policySetList, options.policySetName)

	// Check if workspace exists and got its ID
	workspace, err := c.Workspaces.Read(ctx, options.TerraformOrganization, options.workspaceName)
	if err != nil {
		return err
	}

	policySetRemoveWorkspacesOptions := tfe.PolicySetRemoveWorkspacesOptions{
		Workspaces: []*tfe.Workspace{
			{ID: workspace.ID},
		},
	}

	// Remove workspaces from a policy set
	if err := c.PolicySets.RemoveWorkspaces(ctx, policySetID, policySetRemoveWorkspacesOptions); err != nil {
		return err
	}
	fmt.Println("Workspace(s) '" + options.workspaceName + "' detached from policy set '" + options.policySetName + "' successfully!")

	return nil
}
