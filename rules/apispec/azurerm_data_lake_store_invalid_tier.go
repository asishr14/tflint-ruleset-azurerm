// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataLakeStoreInvalidTierRule checks the pattern is valid
type AzurermDataLakeStoreInvalidTierRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDataLakeStoreInvalidTierRule returns new rule with default attributes
func NewAzurermDataLakeStoreInvalidTierRule() *AzurermDataLakeStoreInvalidTierRule {
	return &AzurermDataLakeStoreInvalidTierRule{
		resourceType:  "azurerm_data_lake_store",
		attributeName: "tier",
		enum: []string{
			"Consumption",
			"Commitment_1TB",
			"Commitment_10TB",
			"Commitment_100TB",
			"Commitment_500TB",
			"Commitment_1PB",
			"Commitment_5PB",
		},
	}
}

// Name returns the rule name
func (r *AzurermDataLakeStoreInvalidTierRule) Name() string {
	return "azurerm_data_lake_store_invalid_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataLakeStoreInvalidTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataLakeStoreInvalidTierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataLakeStoreInvalidTierRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataLakeStoreInvalidTierRule) Check(runner tflint.Runner) error {
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
