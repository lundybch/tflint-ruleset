package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type VariableDefaultRule struct{}

func NewVariableDefaultRule() *VariableDefaultRule {
	return &VariableDefaultRule{}
}

func (r *VariableDefaultRule) Name() string {
	return "custom_variable_default"
}

func (r *VariableDefaultRule) Enabled() bool {
	return true
}

func (r *VariableDefaultRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *VariableDefaultRule) Link() string {
	return ""
}

func (r *VariableDefaultRule) Metadata() *tflint.RuleMetadata {
	return &tflint.RuleMetadata{
		Version: "0.1.0",
	}
}

func (r *VariableDefaultRule) Check(runner tflint.Runner) error {
	return runner.WalkTerraformBlocks("variable", func(block *tflint.Block) error {
		body, _, diag := runner.GetBodyContent(block.Body, nil)
		if diag.HasErrors() {
			return diag
		}

		// Check for the "default" attribute
		attr, exists := body.Attributes["default"]
		if !exists {
			runner.EmitIssue(
				r,
				"Terraform variable is missing a default value",
				block.DefRange,
			)
		}

		return nil
	})
}
