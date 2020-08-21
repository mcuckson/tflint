package terraformrules

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_TerraformResourcesHaveRequiredProviders(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{

		{
			Name: "version set",
			Content: `
terraform {
  required_providers {
	azuread = "1.0"
  }
}


resource "azuread_application" "example" {
 name = "ExampleApp"
}

`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewTerraformResourcesHaveRequiredProvidersRule()

	for _, tc := range cases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			runner := tflint.TestRunner(t, map[string]string{"module.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			tflint.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
