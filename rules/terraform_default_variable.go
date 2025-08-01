package rules

import (
  "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type VariableDefaultRule struct{}

func NewVariableDefaultRule() *VariableDefaultRule {
  return &VariableDefaultRule{}
}

func (r *VariableDefaultRule) Name() string { return "custom_variable_default" }
func (r *VariableDefaultRule) Enabled() bool { return true }
func (r *VariableDefaultRule) Severity() tflint.Severity { return tflint.ERROR }
func (r *VariableDefaultRule) Link() string { return "" }

func (r *VariableDefaultRule) Check(runner tflint.Runner) error {
  return runner.WalkBlocks("variable", func(block *tflint.Block) error {
    if _, exists := block.Body.Attributes["default"]; !exists {
      return runner.EmitIssue(r, "Variable should have a default value", block.DefRange)
    }
    return nil
  })
}
