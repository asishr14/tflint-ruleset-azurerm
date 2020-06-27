// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule checks the pattern is valid
type AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualMachineScaleSetInvalidEvictionPolicyRule returns new rule with default attributes
func NewAzurermVirtualMachineScaleSetInvalidEvictionPolicyRule() *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule {
	return &AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule{
		resourceType:  "azurerm_virtual_machine_scale_set",
		attributeName: "eviction_policy",
		enum: []string{
			"Deallocate",
			"Delete",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Name() string {
	return "azurerm_virtual_machine_scale_set_invalid_eviction_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as eviction_policy`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
