package terraformrules

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// TerraformResourcesHaveRequiredProviders checks whether Terraform sets version constraints for all declared resources
type TerraformResourcesHaveRequiredProviders struct{}

// NewTerraformResourcesHaveRequiredProviders returns new rule with default attributes
func NewTerraformResourcesHaveRequiredProviders() *TerraformResourcesHaveRequiredProviders {
	return &TerraformResourcesHaveRequiredProviders{}
}

// Name returns the rule name
func (r *TerraformResourcesHaveRequiredProviders) Name() string {
	return "terraform_resources_have_required_providers"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformResourcesHaveRequiredProviders) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *TerraformResourcesHaveRequiredProviders) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *TerraformResourcesHaveRequiredProviders) Link() string {
	return tflint.ReferenceLink(r.Name())
}

// Check checks whether declared resources have valid version constraints
func (r *TerraformResourcesHaveRequiredProviders) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	resources := make(map[string]hcl.Range)
	module := runner.TFConfig.Module

	for _, res := range module.ManagedResources {
		if _, ok := resources[res.Provider.Type]; !ok {
			resources[res.Provider.Type] = res.DeclRange
		}
	}

	for name, decl := range resources {
		if _, ok := module.ProviderRequirements.RequiredProviders[name]; !ok {
			runner.EmitIssue(r, fmt.Sprintf(`Missing version constraint for provider "%s" in "required_providers"`, name), decl)
		}
	}

	return nil
}
