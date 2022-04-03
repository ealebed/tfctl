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

// saveOptions represents options for save command
type saveOptions struct {
	*policySetOptions
	policySetName string
	global        bool
	policiesPath  string
	repoName      string
	repoBranch    string
	tokenID       string
}

// NewPolicySetSaveCmd returns new policy set save command
func NewPolicySetSaveCmd(policySetOptions *policySetOptions) *cobra.Command {
	options := &saveOptions{
		policySetOptions: policySetOptions,
	}

	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{"create"},
		Short:   "create a policy set and associate it with terraform organization",
		Long:    "create a policy set (or update if already exists) and associate it with terraform organization",
		Example: "tfctl policySet save ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return savePolicySet(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.policySetName, "policySet", "p", "", "terraform policy set name for creation")
	cmd.Flags().BoolVar(&options.global, "global", false, "Optional: Whether or not the policy set is global")
	cmd.Flags().StringVar(&options.policiesPath, "policiesPath", "/", "Optional: The sub-path within the attached VCS repository to ingress. All files and directories outside of this sub-path will be ignored. This option may only be specified when a VCS repo is present.")

	cmd.Flags().StringVar(&options.repoName, "repoName", "", "full VCS repository identifier (with user or organization path), e.g. 'ealebed/gcp-sentinel-policies'")
	cmd.Flags().StringVar(&options.repoBranch, "repoBranch", "master", "VCS repository branch name to read policies from")
	cmd.Flags().StringVar(&options.tokenID, "tokenID", "", "terraform OAuth token ID. Can be obtained from output 'tfctl OAuthClient get [--providerType=...]'")

	cmd.MarkFlagRequired("policySet")
	cmd.MarkFlagRequired("repoName")
	cmd.MarkFlagRequired("tokenID")

	return cmd
}

func savePolicySet(cmd *cobra.Command, options *saveOptions) error {
	c := options.TClient
	ctx := context.Background()

	var policySet *tfe.PolicySet

	// Fill needed create options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#PolicySetCreateOptions
	policySetCreateOptions := tfe.PolicySetCreateOptions{
		Name:         tfe.String(options.policySetName),
		Global:       tfe.Bool(options.global),
		PoliciesPath: tfe.String(options.policiesPath),
		VCSRepo: &tfe.VCSRepoOptions{
			Branch:            tfe.String(options.repoBranch),
			Identifier:        tfe.String(options.repoName),
			IngressSubmodules: tfe.Bool(false),
			OAuthTokenID:      tfe.String(options.tokenID),
		},
	}

	// Fill needed update options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#PolicySetUpdateOptions
	policySetUpdateOptions := tfe.PolicySetUpdateOptions{
		Name:         tfe.String(options.policySetName),
		Global:       tfe.Bool(options.global),
		PoliciesPath: tfe.String(options.policiesPath),
		VCSRepo: &tfe.VCSRepoOptions{
			Branch:            tfe.String(options.repoBranch),
			Identifier:        tfe.String(options.repoName),
			IngressSubmodules: tfe.Bool(false),
			OAuthTokenID:      tfe.String(options.tokenID),
		},
	}

	// List all the policy sets for a given organization and filter policy set ID by provided name
	policySetList, err := c.PolicySets.List(ctx, options.TerraformOrganization, &tfe.PolicySetListOptions{})
	if err != nil {
		return err
	}
	policySetID := utils.GetPolicySetID(policySetList, options.policySetName)

	if policySetID == "" {
		// Create a policy set and associate it with an organization
		policySet, err = c.PolicySets.Create(ctx, options.TerraformOrganization, policySetCreateOptions)
		if err != nil {
			return err
		}
	} else {
		// Update an existing policy set
		policySet, err = c.PolicySets.Update(ctx, policySetID, policySetUpdateOptions)
		if err != nil {
			return err
		}
	}

	output.JsonPrettyOutput(policySet, "policySet")

	return nil
}
