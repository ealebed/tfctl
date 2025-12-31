# Terraform Cloud/Enterprise command line tool

This is a CLI to help changing and doing stuff in Terraform Cloud.

## Get
```bash
git clone git@github.com:ealebed/tfctl.git
```

## Build
```bash
cd tfctl
```

```bash
go install github.com/ealebed/tfctl
```

or
```bash
make install
```

## Set your env-variables
```bash
export TF_ORG=
export TF_TOKEN=
```

or execute `terraform login` to generate valid 'credentials.tfrc.json' in user home directory.
[More info](https://www.terraform.io/cli/commands/login) about login.

## Use

```bash
tfctl -h
```
---

## Syntax

Use the following syntax to run `tfctl` commands from your terminal window:

```bash
tfctl [command] [flags]
```
---

### Flags:

| Flag      | Description |
| --------- | ----------- |
|  -x, --expand       |  Expand output with all possible values
|  -h, --help         |  help for this command
|      --host string  |  Terraform Enterprise (Cloud) host (default "app.terraform.io")
|      --org string   |  Terraform Enterprise (Cloud) organization name
|      --token string |  Terraform Enterprise (Cloud) token
|  -v, --version      |  version for this command

### Available Commands:

| Command   | Description |
| --------- | ----------- |
|  OAuthClient | Work with terraform OAuth clients
|  completion  | Generate the autocompletion script for the specified shell
|  help        | Help about any command
|  policySet   | Work with terraform policy sets
|  variable    | Work with terraform variables
|  ws          | Work with terraform workspaces

### OAuthClient Subcommands:

| Subcommand   | Description |
| --------- | ----------- |
|  delete | Delete an OAuth client
|  get  | Read an OAuth client
|  list | List all the OAuth clients for a given organization
|  save  | Create an OAuth client to connect an organization and a VCS provider

### policySet Subcommands:

| Subcommand   | Description |
| --------- | ----------- |
|  attach | Attach workspace(s) to a policy set in terraform organization
|  delete  | Delete a policy set from provided organization
|  detach | Remove workspace(s) from a policy set in terraform organization
|  get  | Read a policy set by its ID from provided terraform organization
|  list | List all the policy sets for a given organization
|  save  | Create a policy set and associate it with terraform organization

### variable Subcommands:

| Subcommand   | Description |
| --------- | ----------- |
|  delete | Delete variable from provided terraform workspace
|  get  | Read a variable by its ID from provided terraform workspace
|  list | List all the variables associated with the given workspace
|  save  | Save (create) variable in provided terraform workspace

### ws (workspace) Subcommands:

| Subcommand   | Description |
| --------- | ----------- |
|  delete | Delete a terraform workspace by its name
|  get  | Read a workspace by its name and organization name
|  list | List all the workspaces within an organization
|  save  | Save (create) given terraform workspace

## Examples: Common operations

### Manage OAuth clients

```bash
# Create a new OAuth client for GitHub service provider
tfctl OAuthClient save --providerType=github --token=ghp_3gtKvymPvIpx3Of34dfgRCRSOdY1ApR3rpeTW

# Get expanded [-x] information about OAuth client from GitLab service provider
tfctl OAuthClient get --providerType gitlab -x

# List all the OAuth clients in 'ealebed' terraform organization
tfctl OAuthClient list --org ealebed

# Delete OAuth client from GitLab service provider
tfctl OAuthClient delete --providerType gitlab
```

### Manage policy sets

```bash
# Add workspace 'vertex-ai-notebooks' to policy set 'test-policy-set'
tfctl policySet attach -p test-policy-set -w vertex-ai-notebooks

# Get expanded [-x] information about policy set 'test-policy-set'
tfctl policySet get -p test-policy-set -x

# Remove workspace 'vertex-ai-notebooks' from policy set 'test-policy-set'
tfctl policySet detach -p test-policy-set -w vertex-ai-notebooks

# Delete policy set 'test-policy-set'
tfctl policySet delete --policySet=test-policy-set

# List all policy sets in 'ealebed' terraform organization
tfctl policySet list --org ealebed

# Save policy set 'test-gh-policy-set' from GitHub repository 'ealebed/sentinel-policies'
tfctl policySet save -p test-gh-policy-set --repoName ealebed/sentinel-policies --tokenID ot-6PdBa6bXPWeyGZBm
```

### Manage variables

```bash
# Create a new sensitive variable 'foo' with value 'bar' in 'gitlab-tfc-demo' terraform workspace
tfctl variable save -w gitlab-tfc-demo --key foo --value bar --sensitive

# Get expanded [-x] information about variable 'varSet' from 'gitlab-tfc-demo' terraform workspace
tfctl variable get -w gitlab-tfc-demo -v varSet

# List all variables in 'gitlab-tfc-demo' terraform workspace
tfctl variable list -w gitlab-tfc-demo

# Delete variable 'testVar' from 'gitlab-tfc-demo' terraform workspace
tfctl variable delete -w gitlab-tfc-demo -v testVar
```

### Manage workspaces

```bash
# Create a new 'gitlab-tfc-demo' terraform workspace
tfctl ws save -w gitlab-tfc-demo

# Get expanded [-x] information about 'gitlab-tfc-demo' terraform workspace
tfctl ws get -w gitlab-tfc-demo

# List all terraform workspaces
tfctl ws list

# Delete 'gitlab-tfc-demo' terraform workspace
tfctl ws delete -w gitlab-tfc-demo
```

TODO:
- Add global '--dry-run' flag for destructive operations
- Add colored output
