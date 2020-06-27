// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetappVolumeInvalidNameRule checks the pattern is valid
type AzurermNetappVolumeInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermNetappVolumeInvalidNameRule returns new rule with default attributes
func NewAzurermNetappVolumeInvalidNameRule() *AzurermNetappVolumeInvalidNameRule {
	return &AzurermNetappVolumeInvalidNameRule{
		resourceType:  "azurerm_netapp_volume",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`),
	}
}

// Name returns the rule name
func (r *AzurermNetappVolumeInvalidNameRule) Name() string {
	return "azurerm_netapp_volume_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetappVolumeInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetappVolumeInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetappVolumeInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetappVolumeInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][a-zA-Z0-9\-_]{0,63}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
