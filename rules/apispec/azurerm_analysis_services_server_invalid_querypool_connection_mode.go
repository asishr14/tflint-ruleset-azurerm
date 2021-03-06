// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule checks the pattern is valid
type AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule returns new rule with default attributes
func NewAzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule() *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule {
	return &AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule{
		resourceType:  "azurerm_analysis_services_server",
		attributeName: "querypool_connection_mode",
		enum: []string{
			"All",
			"ReadOnly",
		},
	}
}

// Name returns the rule name
func (r *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule) Name() string {
	return "azurerm_analysis_services_server_invalid_querypool_connection_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as querypool_connection_mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
