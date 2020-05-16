// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermServicebusTopicInvalidStatusRule checks the pattern is valid
type AzurermServicebusTopicInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServicebusTopicInvalidStatusRule returns new rule with default attributes
func NewAzurermServicebusTopicInvalidStatusRule() *AzurermServicebusTopicInvalidStatusRule {
	return &AzurermServicebusTopicInvalidStatusRule{
		resourceType:  "azurerm_servicebus_topic",
		attributeName: "status",
		enum: []string{
			"Active",
			"Disabled",
			"Restoring",
			"SendDisabled",
			"ReceiveDisabled",
			"Creating",
			"Deleting",
			"Renaming",
			"Unknown",
		},
	}
}

// Name returns the rule name
func (r *AzurermServicebusTopicInvalidStatusRule) Name() string {
	return "azurerm_servicebus_topic_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServicebusTopicInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServicebusTopicInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServicebusTopicInvalidStatusRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermServicebusTopicInvalidStatusRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}