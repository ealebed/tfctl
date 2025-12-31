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

//nolint:dupl // This file has similar command setup pattern to get.go
package workspace

import (
	"context"

	"github.com/ealebed/tfctl/pkg/output"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

// saveOptions represents options for save command
type saveOptions struct {
	*workspaceOptions
	workspaceName string
}

// NewWorkspaceSaveCmd returns new workspace save command
func NewWorkspaceSaveCmd(workspaceOptions *workspaceOptions) *cobra.Command {
	options := &saveOptions{
		workspaceOptions: workspaceOptions,
	}

	cmd := &cobra.Command{
		Use:     "save",
		Aliases: []string{"create"},
		Short:   "save (create) given terraform workspace",
		Long:    "save (create or update if already exists) given terraform workspace",
		Example: "tfctl ws save [--workspace=...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return createWorkspace(cmd, options)
		},
	}

	cmd.Flags().StringVarP(&options.workspaceName, "workspace", "w", "", "name terraform workspace to create")
	if err := cmd.MarkFlagRequired("workspace"); err != nil {
		return nil
	}

	return cmd
}

func createWorkspace(_ *cobra.Command, options *saveOptions) error {
	c := options.TClient
	ctx := context.Background()

	var workspace *tfe.Workspace
	var err error

	// Fill needed create options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#WorkspaceCreateOptions
	createOptions := tfe.WorkspaceCreateOptions{
		Name:        tfe.String(options.workspaceName),
		Description: tfe.String("Description for workspace " + options.workspaceName),
		Tags: []*tfe.Tag{
			{Name: options.workspaceName},
		},
	}

	// Fill needed update options from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#WorkspaceUpdateOptions
	updateOptions := tfe.WorkspaceUpdateOptions{
		Name:        tfe.String(options.workspaceName),
		Description: tfe.String("Description for workspace " + options.workspaceName),
	}

	existingWorkspace, _ := c.Workspaces.Read(ctx, options.TerraformOrganization, options.workspaceName)
	if existingWorkspace != nil {
		// Update settings of an existing workspace
		workspace, err = c.Workspaces.Update(ctx, options.TerraformOrganization, options.workspaceName, updateOptions)
		if err != nil {
			return err
		}
	} else {
		// Create a new workspace
		workspace, err = c.Workspaces.Create(ctx, options.TerraformOrganization, createOptions)
		if err != nil {
			return err
		}
	}

	output.JsonPrettyOutput(workspace, "workspace")

	return nil
}
