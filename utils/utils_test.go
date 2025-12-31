package utils

import (
	"testing"

	"github.com/hashicorp/go-tfe"
)

func TestGetVariableID(t *testing.T) {
	tests := []struct {
		name         string
		variables    *tfe.VariableList
		variableName string
		want         string
	}{
		{
			name: "variable found",
			variables: &tfe.VariableList{
				Items: []*tfe.Variable{
					{ID: "var-1", Key: "test_key"},
					{ID: "var-2", Key: "another_key"},
					{ID: "var-3", Key: "third_key"},
				},
			},
			variableName: "another_key",
			want:         "var-2",
		},
		{
			name: "variable not found",
			variables: &tfe.VariableList{
				Items: []*tfe.Variable{
					{ID: "var-1", Key: "test_key"},
					{ID: "var-2", Key: "another_key"},
				},
			},
			variableName: "nonexistent_key",
			want:         "",
		},
		{
			name:         "empty variable list",
			variables:    &tfe.VariableList{Items: []*tfe.Variable{}},
			variableName: "any_key",
			want:         "",
		},
		{
			name:         "nil variable list",
			variables:    nil,
			variableName: "any_key",
			want:         "",
		},
		{
			name: "case sensitive match",
			variables: &tfe.VariableList{
				Items: []*tfe.Variable{
					{ID: "var-1", Key: "TestKey"},
					{ID: "var-2", Key: "testkey"},
				},
			},
			variableName: "TestKey",
			want:         "var-1",
		},
		{
			name: "first match returned",
			variables: &tfe.VariableList{
				Items: []*tfe.Variable{
					{ID: "var-1", Key: "duplicate"},
					{ID: "var-2", Key: "duplicate"},
				},
			},
			variableName: "duplicate",
			want:         "var-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetVariableID(tt.variables, tt.variableName)
			if got != tt.want {
				t.Errorf("GetVariableID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPolicySetID(t *testing.T) {
	tests := []struct {
		name          string
		policySets    *tfe.PolicySetList
		policySetName string
		want          string
	}{
		{
			name: "policy set found",
			policySets: &tfe.PolicySetList{
				Items: []*tfe.PolicySet{
					{ID: "ps-1", Name: "test-policy"},
					{ID: "ps-2", Name: "another-policy"},
					{ID: "ps-3", Name: "third-policy"},
				},
			},
			policySetName: "another-policy",
			want:          "ps-2",
		},
		{
			name: "policy set not found",
			policySets: &tfe.PolicySetList{
				Items: []*tfe.PolicySet{
					{ID: "ps-1", Name: "test-policy"},
					{ID: "ps-2", Name: "another-policy"},
				},
			},
			policySetName: "nonexistent-policy",
			want:          "",
		},
		{
			name:          "empty policy set list",
			policySets:    &tfe.PolicySetList{Items: []*tfe.PolicySet{}},
			policySetName: "any-policy",
			want:          "",
		},
		{
			name:          "nil policy set list",
			policySets:    nil,
			policySetName: "any-policy",
			want:          "",
		},
		{
			name: "case sensitive match",
			policySets: &tfe.PolicySetList{
				Items: []*tfe.PolicySet{
					{ID: "ps-1", Name: "TestPolicy"},
					{ID: "ps-2", Name: "testpolicy"},
				},
			},
			policySetName: "TestPolicy",
			want:          "ps-1",
		},
		{
			name: "first match returned",
			policySets: &tfe.PolicySetList{
				Items: []*tfe.PolicySet{
					{ID: "ps-1", Name: "duplicate"},
					{ID: "ps-2", Name: "duplicate"},
				},
			},
			policySetName: "duplicate",
			want:          "ps-1",
		},
		{
			name: "empty policy set name",
			policySets: &tfe.PolicySetList{
				Items: []*tfe.PolicySet{
					{ID: "ps-1", Name: ""},
					{ID: "ps-2", Name: "non-empty"},
				},
			},
			policySetName: "",
			want:          "ps-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPolicySetID(tt.policySets, tt.policySetName)
			if got != tt.want {
				t.Errorf("GetPolicySetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOAuthClientID(t *testing.T) {
	tests := []struct {
		name                string
		oauthClient         *tfe.OAuthClientList
		serviceProviderName string
		want                string
		wantError           bool
	}{
		{
			name: "exact match found",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github"},
					{ID: "oauth-2", ServiceProviderName: "gitlab"},
					{ID: "oauth-3", ServiceProviderName: "bitbucket"},
				},
			},
			serviceProviderName: "github",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "case insensitive match",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "GitHub"},
					{ID: "oauth-2", ServiceProviderName: "GitLab"},
				},
			},
			serviceProviderName: "github",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "case insensitive match reverse",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github"},
					{ID: "oauth-2", ServiceProviderName: "gitlab"},
				},
			},
			serviceProviderName: "GitHub",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "partial match",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github_enterprise"},
					{ID: "oauth-2", ServiceProviderName: "gitlab_hosted"},
				},
			},
			serviceProviderName: "github",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "not found",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github"},
					{ID: "oauth-2", ServiceProviderName: "gitlab"},
				},
			},
			serviceProviderName: "bitbucket",
			want:                "",
			wantError:           false,
		},
		{
			name:                "empty oauth client list",
			oauthClient:         &tfe.OAuthClientList{Items: []*tfe.OAuthClient{}},
			serviceProviderName: "github",
			want:                "",
			wantError:           false,
		},
		{
			name:                "nil oauth client list",
			oauthClient:         nil,
			serviceProviderName: "github",
			want:                "",
			wantError:           false,
		},
		{
			name: "first match returned",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github"},
					{ID: "oauth-2", ServiceProviderName: "github"},
				},
			},
			serviceProviderName: "github",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "empty service provider name",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: ""},
					{ID: "oauth-2", ServiceProviderName: "github"},
				},
			},
			serviceProviderName: "",
			want:                "oauth-1",
			wantError:           false,
		},
		{
			name: "special characters in service provider name",
			oauthClient: &tfe.OAuthClientList{
				Items: []*tfe.OAuthClient{
					{ID: "oauth-1", ServiceProviderName: "github-enterprise"},
					{ID: "oauth-2", ServiceProviderName: "gitlab_hosted"},
				},
			},
			serviceProviderName: "github-enterprise",
			want:                "oauth-1",
			wantError:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOAuthClientID(tt.oauthClient, tt.serviceProviderName)
			if got != tt.want {
				t.Errorf("GetOAuthClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}
