// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetappPoolInvalidServiceLevelRule checks the pattern is valid
type AzurermNetappPoolInvalidServiceLevelRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermNetappPoolInvalidServiceLevelRule returns new rule with default attributes
func NewAzurermNetappPoolInvalidServiceLevelRule() *AzurermNetappPoolInvalidServiceLevelRule {
	return &AzurermNetappPoolInvalidServiceLevelRule{
		resourceType:  "azurerm_netapp_pool",
		attributeName: "service_level",
		enum: []string{
			"Standard",
			"Premium",
			"Ultra",
		},
	}
}

// Name returns the rule name
func (r *AzurermNetappPoolInvalidServiceLevelRule) Name() string {
	return "azurerm_netapp_pool_invalid_service_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetappPoolInvalidServiceLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetappPoolInvalidServiceLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetappPoolInvalidServiceLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetappPoolInvalidServiceLevelRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as service_level`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
