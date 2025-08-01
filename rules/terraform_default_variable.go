package rules

import (
  "fmt"
  "github.com/terraform-linters/tflint-plugin-sdk/tflint"
  "github.com/zclconf/go-cty/cty"
  "github.com/hashicorp/hcl/v2"
)

type VariableHasDefaultRule struct {
  tflint.DefaultRule
}

func (r *VariableHasDefaultRule) Name() string     { return "custom_variable_has_default" }
func (r *VariableHasDefaultRule) Enabled() bool    { return true }
func (r *VariableHasDefaultRule) Severity() string { return tflint.ERROR }
func (r *VariableHasDefaultRule) Link() string     { return "" }

func (r *VariableHasDefaultRule) Check(runner tflint.Runner) error {
  modules, err := runner.GetModuleContent(&tflint.BodySchema{
    Blocks: []tflint.BlockSchema{
      {
        Type:       "variable",
        LabelNames: []string{"name"},
        Body: &tflint.BodySchema{
          Attributes: []tflint.AttributeSchema{{Name: "default"}},
        },
      },
    },
  }, nil)
  if err != nil {
    return err
  }

  for _, module := range modules {
    for _, block := range module.Blocks {
      varName := block.Labels[0]
      _, hasDefault := block.Body.Attributes["default"]
      if !hasDefault {
        runner.EmitIssue(r,
          fmt.Sprintf("variable %q has no default value", varName),
          block.Body.Range,
        )
      }
    }
  }
  return nil
}