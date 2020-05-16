<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_eventgrid_event_subscription_invalid_event_delivery_schema

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- EventGridSchema
- CloudEventV01Schema
- CustomInputSchema

## Example

```hcl
resource "azurerm_eventgrid_event_subscription" "foo" {
  event_delivery_schema = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as event_delivery_schema (azurerm_eventgrid_event_subscription_invalid_event_delivery_schema)

  on template.tf line 2:
  2:   event_delivery_schema = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_eventgrid_event_subscription_invalid_event_delivery_schema.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2019-02-01-preview/EventGrid.json