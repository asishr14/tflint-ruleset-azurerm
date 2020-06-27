// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermHdinsightMlServicesClusterInvalidTierRule checks the pattern is valid
type AzurermHdinsightMlServicesClusterInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermHdinsightMlServicesClusterInvalidTierRule returns new rule with default attributes
func NewAzurermHdinsightMlServicesClusterInvalidTierRule() *AzurermHdinsightMlServicesClusterInvalidTierRule {
	return &AzurermHdinsightMlServicesClusterInvalidTierRule{
		resourceType:  "azurerm_hdinsight_ml_services_cluster",
		attributeName: "tier",
		enum: []string{
			"Standard",
			"Premium",
		},
	}
}

// Name returns the rule name
func (r *AzurermHdinsightMlServicesClusterInvalidTierRule) Name() string {
	return "azurerm_hdinsight_ml_services_cluster_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermHdinsightMlServicesClusterInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermHdinsightMlServicesClusterInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermHdinsightMlServicesClusterInvalidTierRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermHdinsightMlServicesClusterInvalidTierRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as tier`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
