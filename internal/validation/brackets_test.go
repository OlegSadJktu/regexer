package validation

import "testing"

type input struct {
  str string
  result bool
}

func TestBrackets(t *testing.T) {
  inputs := []input{
    {"123[]{}", true},
    {"123[}[}", false},
    {"123[[[{{*(123)}}]]]{}", true},
    {"123[[[{{*(123)233()}}{{}}]]]{}", true},
    {"12[{{[[[{{*(123)233()}}{{}}]]]{}{}}]", false},
  }
  for _, inp := range inputs {
    if CheckBrackets(inp.str) != inp.result {
      t.Errorf("input: %v, expected %v", inp.str, inp.result)
    }
  }

}
