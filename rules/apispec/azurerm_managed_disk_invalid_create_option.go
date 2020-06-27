// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermManagedDiskInvalidCreateOptionRule checks the pattern is valid
type AzurermManagedDiskInvalidCreateOptionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermManagedDiskInvalidCreateOptionRule returns new rule with default attributes
func NewAzurermManagedDiskInvalidCreateOptionRule() *AzurermManagedDiskInvalidCreateOptionRule {
	return &AzurermManagedDiskInvalidCreateOptionRule{
		resourceType:  "azurerm_managed_disk",
		attributeName: "create_option",
		enum: []string{
			"Empty",
			"Attach",
			"FromImage",
			"Import",
			"Copy",
			"Restore",
			"Upload",
		},
	}
}

// Name returns the rule name
func (r *AzurermManagedDiskInvalidCreateOptionRule) Name() string {
	return "azurerm_managed_disk_invalid_create_option"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermManagedDiskInvalidCreateOptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermManagedDiskInvalidCreateOptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermManagedDiskInvalidCreateOptionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermManagedDiskInvalidCreateOptionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as create_option`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
