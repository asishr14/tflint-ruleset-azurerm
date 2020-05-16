// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule checks the pattern is valid
type AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule returns new rule with default attributes
func NewAzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule() *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule {
	return &AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule{
		resourceType:  "azurerm_cost_management_export_resource_group",
		attributeName: "recurrence_type",
		enum: []string{
			"Daily",
			"Weekly",
			"Monthly",
			"Annually",
		},
	}
}

// Name returns the rule name
func (r *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule) Name() string {
	return "azurerm_cost_management_export_resource_group_invalid_recurrence_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCostManagementExportResourceGroupInvalidRecurrenceTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as recurrence_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}