// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermContainerRegistryWebhookInvalidStatusRule checks the pattern is valid
type AzurermContainerRegistryWebhookInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermContainerRegistryWebhookInvalidStatusRule returns new rule with default attributes
func NewAzurermContainerRegistryWebhookInvalidStatusRule() *AzurermContainerRegistryWebhookInvalidStatusRule {
	return &AzurermContainerRegistryWebhookInvalidStatusRule{
		resourceType:  "azurerm_container_registry_webhook",
		attributeName: "status",
		enum: []string{
			"enabled",
			"disabled",
		},
	}
}

// Name returns the rule name
func (r *AzurermContainerRegistryWebhookInvalidStatusRule) Name() string {
	return "azurerm_container_registry_webhook_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermContainerRegistryWebhookInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermContainerRegistryWebhookInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermContainerRegistryWebhookInvalidStatusRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermContainerRegistryWebhookInvalidStatusRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
