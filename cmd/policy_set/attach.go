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

// attachOptions represents options for attach command
type attachOptions struct {
	*policySetOptions
	policySetName string
	workspaceName string
}

// NewPolicySetAttachCmd returns new policy set attach command
func NewPolicySetAttachCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &attachOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "attach",
		Short:   "attach workspace(s) to a policy set in terraform organization",
		Long:    "attach workspace(s) to a policy set in terraform organization",
		Example: "tfctl policySet attach [--policySet=...] [--name=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return attachToPolicySet(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.policySetName, "policySet", "p", "", "terraform policy set name")
	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "terraform policy workspace name for attaching to a policy set")
	cmd.MarkFlagRequired("policySet")
	cmd.MarkFlagRequired("workspace")

	return cmd
}

func attachToPolicySet(cmd *cobra.Command, options *attachOptions) error {
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

	policySetAddWorkspacesOptions := tfe.PolicySetAddWorkspacesOptions{
		Workspaces: []*tfe.Workspace{
			{ID: workspace.ID},
		},
	}

	// Add workspaces to a policy set
	if err := c.PolicySets.AddWorkspaces(ctx, policySetID, policySetAddWorkspacesOptions); err != nil {
		return err
	} else {
		fmt.Println("Workspace(s) '" + options.workspaceName + "' attached to policy set '" + options.policySetName + "' successfully!")
	}

	return nil
}
