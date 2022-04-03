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

package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/go-tfe"
)

// Take needed fields from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#OAuthToken
type outputOAuthToken struct {
	ID string `jsonapi:"primary,oauth-tokens"`
}

// Take needed fields from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#OAuthClient
type outputOAuthClient struct {
	ID                  string                  `jsonapi:"primary,oauth-clients"`
	APIURL              string                  `jsonapi:"attr,api-url"`
	HTTPURL             string                  `jsonapi:"attr,http-url"`
	ServiceProvider     tfe.ServiceProviderType `jsonapi:"attr,service-provider"`
	ServiceProviderName string                  `jsonapi:"attr,service-provider-display-name"`
	OAuthTokens         []*outputOAuthToken     `jsonapi:"relation,oauth-tokens"`
}

// Take needed fields from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#PolicySet
type outputPolicySet struct {
	ID             string       `jsonapi:"primary,policy-sets"`
	Name           string       `jsonapi:"attr,name"`
	Description    string       `jsonapi:"attr,description"`
	Global         bool         `jsonapi:"attr,global"`
	PoliciesPath   string       `jsonapi:"attr,policies-path"`
	VCSRepo        *tfe.VCSRepo `jsonapi:"attr,vcs-repo"`
	WorkspaceCount int          `jsonapi:"attr,workspace-count"`
}

// Take needed fields from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#Variable
type outputVariable struct {
	ID          string           `jsonapi:"primary,vars"`
	Key         string           `jsonapi:"attr,key"`
	Value       string           `jsonapi:"attr,value"`
	Description string           `jsonapi:"attr,description"`
	Category    tfe.CategoryType `jsonapi:"attr,category"`
	HCL         bool             `jsonapi:"attr,hcl"`
	Sensitive   bool             `jsonapi:"attr,sensitive"`
}

// Take needed fields from https://pkg.go.dev/github.com/hashicorp/go-tfe@v1.1.0#Workspace
type outputWorkspace struct {
	ID               string   `jsonapi:"primary,workspaces"`
	Description      string   `jsonapi:"attr,description"`
	ExecutionMode    string   `jsonapi:"attr,execution-mode"`
	Name             string   `jsonapi:"attr,name"`
	TerraformVersion string   `jsonapi:"attr,terraform-version"`
	TagNames         []string `jsonapi:"attr,tag-names"`
}

func marshalToJson(input interface{}) ([]byte, error) {
	pretty, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to json: %v", err)
	}
	return pretty, nil
}

func JsonOutput(input interface{}) {
	res, err := marshalToJson(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(res))
}

func JsonPrettyOutput(input interface{}, inputType string) {
	tmp, _ := json.MarshalIndent(input, "", " ")

	switch inputType {
	case "workspace":
		var out *outputWorkspace
		json.Unmarshal(tmp, &out)
		JsonOutput(out)
	case "variable":
		var out *outputVariable
		json.Unmarshal(tmp, &out)
		JsonOutput(out)
	case "policySet":
		var out *outputPolicySet
		json.Unmarshal(tmp, &out)
		JsonOutput(out)
	case "OAuthClient":
		var out *outputOAuthClient
		json.Unmarshal(tmp, &out)
		JsonOutput(out)
	default:
		fmt.Println("ERROR: No such type '" + inputType + "' for output...")
		os.Exit(1)
	}
}
