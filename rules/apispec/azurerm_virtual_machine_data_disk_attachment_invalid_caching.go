// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule checks the pattern is valid
type AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualMachineDataDiskAttachmentInvalidCachingRule returns new rule with default attributes
func NewAzurermVirtualMachineDataDiskAttachmentInvalidCachingRule() *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule {
	return &AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule{
		resourceType:  "azurerm_virtual_machine_data_disk_attachment",
		attributeName: "caching",
		enum: []string{
			"None",
			"ReadOnly",
			"ReadWrite",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule) Name() string {
	return "azurerm_virtual_machine_data_disk_attachment_invalid_caching"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineDataDiskAttachmentInvalidCachingRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as caching`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
