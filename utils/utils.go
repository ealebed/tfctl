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

package utils

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-tfe"
)

// GetVariableID returns variable ID by given name
func GetVariableID(variables *tfe.VariableList, variableName string) string {
	for _, v := range variables.Items {
		if v.Key == variableName {
			return v.ID
		}
	}

	return ""
}

// GetPolicySetID returns policy set ID by given name
func GetPolicySetID(policySets *tfe.PolicySetList, policySetName string) string {
	for _, p := range policySets.Items {
		if p.Name == policySetName {
			return p.ID
		}
	}

	return ""
}

// GetOAuthClientID returns OAuth client ID by given service provider type
func GetOAuthClientID(OAuthClient *tfe.OAuthClientList, serviceProviderName string) string {
	regexPattern := fmt.Sprintf(`(?i)%s`, serviceProviderName)

	for _, o := range OAuthClient.Items {
		matched, _ := regexp.MatchString(regexPattern, o.ServiceProviderName)
		if matched {
			return o.ID
		}
	}

	return ""
}
