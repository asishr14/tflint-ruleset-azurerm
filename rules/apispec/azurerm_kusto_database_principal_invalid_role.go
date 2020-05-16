// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKustoDatabasePrincipalInvalidRoleRule checks the pattern is valid
type AzurermKustoDatabasePrincipalInvalidRoleRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermKustoDatabasePrincipalInvalidRoleRule returns new rule with default attributes
func NewAzurermKustoDatabasePrincipalInvalidRoleRule() *AzurermKustoDatabasePrincipalInvalidRoleRule {
	return &AzurermKustoDatabasePrincipalInvalidRoleRule{
		resourceType:  "azurerm_kusto_database_principal",
		attributeName: "role",
		enum: []string{
			"Admin",
			"Ingestor",
			"Monitor",
			"User",
			"UnrestrictedViewers",
			"Viewer",
		},
	}
}

// Name returns the rule name
func (r *AzurermKustoDatabasePrincipalInvalidRoleRule) Name() string {
	return "azurerm_kusto_database_principal_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKustoDatabasePrincipalInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKustoDatabasePrincipalInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKustoDatabasePrincipalInvalidRoleRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKustoDatabasePrincipalInvalidRoleRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as role`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}