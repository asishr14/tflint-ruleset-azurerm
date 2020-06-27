// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBotConnectionInvalidBotNameRule checks the pattern is valid
type AzurermBotConnectionInvalidBotNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBotConnectionInvalidBotNameRule returns new rule with default attributes
func NewAzurermBotConnectionInvalidBotNameRule() *AzurermBotConnectionInvalidBotNameRule {
	return &AzurermBotConnectionInvalidBotNameRule{
		resourceType:  "azurerm_bot_connection",
		attributeName: "bot_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
	}
}

// Name returns the rule name
func (r *AzurermBotConnectionInvalidBotNameRule) Name() string {
	return "azurerm_bot_connection_invalid_bot_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBotConnectionInvalidBotNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBotConnectionInvalidBotNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBotConnectionInvalidBotNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBotConnectionInvalidBotNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
