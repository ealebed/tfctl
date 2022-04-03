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

package assembler

import (
	"github.com/ealebed/tfctl/cmd"
	"github.com/ealebed/tfctl/cmd/oauth_client"
	"github.com/ealebed/tfctl/cmd/policy_set"
	"github.com/ealebed/tfctl/cmd/variable"
	"github.com/ealebed/tfctl/cmd/workspace"
	"github.com/spf13/cobra"
)

// AddSubCommands adds all the subcommands to the rootCmd.
// rootOpts are passed through to the subcommands.
func AddSubCommands(rootCmd *cobra.Command, rootOpts *cmd.RootOptions) {
	rootCmd.AddCommand(workspace.NewWorkspaceCmd(rootOpts))
	rootCmd.AddCommand(variable.NewVariableCmd(rootOpts))
	rootCmd.AddCommand(policy_set.NewPolicySetCmd(rootOpts))
	rootCmd.AddCommand(oauth_client.NewOAuthClientCmd(rootOpts))
}
