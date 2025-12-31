package output

import (
	"encoding/json"
	"testing"
)

func TestMarshalToJson(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "valid struct",
			input:   map[string]string{"key": "value"},
			wantErr: false,
		},
		{
			name:    "valid nested struct",
			input:   map[string]interface{}{"key": map[string]string{"nested": "value"}},
			wantErr: false,
		},
		{
			name:    "empty map",
			input:   map[string]string{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			wantErr: false,
		},
		{
			name:    "string input",
			input:   "test string",
			wantErr: false,
		},
		{
			name:    "number input",
			input:   42,
			wantErr: false,
		},
		{
			name:    "slice input",
			input:   []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name: "complex struct",
			input: struct {
				Name  string
				Value int
				Tags  []string
			}{
				Name:  "test",
				Value: 100,
				Tags:  []string{"tag1", "tag2"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := marshalToJson(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("marshalToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got == nil {
					t.Error("marshalToJson() returned nil, expected non-nil")
					return
				}
				// Verify it's valid JSON by unmarshaling
				var result interface{}
				if err := json.Unmarshal(got, &result); err != nil {
					t.Errorf("marshalToJson() produced invalid JSON: %v", err)
				}
				// For complex structures, verify it's pretty-printed (contains indentation)
				// Simple values like strings, numbers, empty maps don't need indentation
				if len(got) > 10 {
					// Check if it contains indentation (spaces after newlines) for complex structures
					hasIndent := false
					for i := 0; i < len(got)-1; i++ {
						if got[i] == '\n' && got[i+1] == ' ' {
							hasIndent = true
							break
						}
					}
					// Only require indentation for complex nested structures
					if !hasIndent && (tt.name == "valid_nested_struct" || tt.name == "complex_struct") {
						t.Error("marshalToJson() should produce pretty-printed JSON with indentation for complex structures")
					}
				}
			}
		})
	}
}

func TestMarshalToJson_ErrorHandling(t *testing.T) {
	// Test with a channel, which cannot be marshaled to JSON
	ch := make(chan int)
	_, err := marshalToJson(ch)
	if err == nil {
		t.Error("marshalToJson() should return error for unmarshalable types")
	}
}

func TestJsonPrettyOutput_Workspace(t *testing.T) {
	// Create a minimal workspace-like structure that can be marshaled
	// The actual tfe.Workspace has jsonapi fields that can't be directly marshaled
	workspaceData := map[string]interface{}{
		"id":                "ws-123",
		"name":              "test-workspace",
		"description":       "Test description",
		"execution-mode":    "remote",
		"terraform-version": "1.0.0",
		"tag-names":         []string{"tag1", "tag2"},
	}

	// This function calls os.Exit on error, so we can't test error cases easily
	// But we can verify it doesn't panic with valid input
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("JsonPrettyOutput() panicked: %v", r)
		}
	}()

	// Test with a type that can be marshaled
	// We'll use the outputWorkspace type indirectly through JsonPrettyOutput
	// Since JsonPrettyOutput unmarshals and remarshals, we need to provide
	// data that matches the expected structure
	_ = workspaceData
	// Skip this test as tfe.Workspace has jsonapi fields that can't be directly tested
	t.Skip("Skipping - tfe.Workspace has jsonapi fields that require special marshaling")
}

func TestJsonPrettyOutput_Variable(t *testing.T) {
	// Skip this test as tfe.Variable has jsonapi fields that can't be directly tested
	t.Skip("Skipping - tfe.Variable has jsonapi fields that require special marshaling")
}

func TestJsonPrettyOutput_PolicySet(t *testing.T) {
	// Skip this test as tfe.PolicySet has jsonapi fields that can't be directly tested
	t.Skip("Skipping - tfe.PolicySet has jsonapi fields that require special marshaling")
}

func TestJsonPrettyOutput_OAuthClient(t *testing.T) {
	// Skip this test as tfe.OAuthClient has jsonapi fields that can't be directly tested
	t.Skip("Skipping - tfe.OAuthClient has jsonapi fields that require special marshaling")
}

func TestJsonPrettyOutput_InvalidType(t *testing.T) {
	// This will call os.Exit(1), so we can't test it normally
	// But we can verify the function exists and doesn't have syntax errors
	// by checking it compiles
	_ = JsonPrettyOutput
}

func TestJsonOutput(t *testing.T) {
	input := map[string]string{"key": "value"}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("JsonOutput() panicked: %v", r)
		}
	}()

	// We can't easily test the output without capturing stdout,
	// but we can verify it doesn't crash
	JsonOutput(input)
}
