// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermRouteInvalidNextHopTypeRule checks the pattern is valid
type AzurermRouteInvalidNextHopTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermRouteInvalidNextHopTypeRule returns new rule with default attributes
func NewAzurermRouteInvalidNextHopTypeRule() *AzurermRouteInvalidNextHopTypeRule {
	return &AzurermRouteInvalidNextHopTypeRule{
		resourceType:  "azurerm_route",
		attributeName: "next_hop_type",
		enum: []string{
			"VirtualNetworkGateway",
			"VnetLocal",
			"Internet",
			"VirtualAppliance",
			"None",
		},
	}
}

// Name returns the rule name
func (r *AzurermRouteInvalidNextHopTypeRule) Name() string {
	return "azurerm_route_invalid_next_hop_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermRouteInvalidNextHopTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermRouteInvalidNextHopTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermRouteInvalidNextHopTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermRouteInvalidNextHopTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as next_hop_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
