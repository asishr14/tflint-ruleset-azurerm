// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule checks the pattern is valid
type AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule returns new rule with default attributes
func NewAzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule() *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule {
	return &AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule{
		resourceType:  "azurerm_traffic_manager_profile",
		attributeName: "traffic_routing_method",
		enum: []string{
			"Performance",
			"Priority",
			"Weighted",
			"Geographic",
			"MultiValue",
			"Subnet",
		},
	}
}

// Name returns the rule name
func (r *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule) Name() string {
	return "azurerm_traffic_manager_profile_invalid_traffic_routing_method"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermTrafficManagerProfileInvalidTrafficRoutingMethodRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as traffic_routing_method`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
