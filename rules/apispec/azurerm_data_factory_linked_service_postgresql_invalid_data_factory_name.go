// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule checks the pattern is valid
type AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule returns new rule with default attributes
func NewAzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule() *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule {
	return &AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule{
		resourceType:  "azurerm_data_factory_linked_service_postgresql",
		attributeName: "data_factory_name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule) Name() string {
	return "azurerm_data_factory_linked_service_postgresql_invalid_data_factory_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryLinkedServicePostgresqlInvalidDataFactoryNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
