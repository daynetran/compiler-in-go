package compiler

import (
	"monkey/code"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}
