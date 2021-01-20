package program

import "github.com/bdlm/errors/v2"

type instructionType string

const (
	// Data stored on ROM
	T_CONST instructionType = "const"

	// Label for JMP instructions.
	T_LABEL instructionType = "label"

	// Instruction
	T_INSTR instructionType = "instr"

	// Subroutine
	T_SUB instructionType = "sub"

	// Subroutine end
	T_SUBEND instructionType = "subend"

	// Variable Pointer
	T_VAR instructionType = "var"

	// Kbit32
	Kbit32 = 32768
)

var (
	errUnknownInstruction = errors.Errorf("unknown instruction type")
	errSyntaxError        = errors.Errorf("invalid syntax")
	errParseError         = errors.Errorf("unknown")
	errCompileError       = errors.Errorf("unknown")
)
