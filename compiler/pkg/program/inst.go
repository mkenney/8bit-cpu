package program

import (
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
)

// Instruction
type Instruction struct {
	Code  string // Line from source file
	Data  int    // Data portion of the instruction code
	Idx   int    // Source file index
	Line  int    // Line number from source file
	Label string

	Inst    int    // Instruction address in compiled source
	Opcodes []byte // Opcodes in compiled source

	Step int // Instruction set step number
	Type instructionType
}

// newInstruction
func newInstruction(code string, idx int) (*Instruction, error) {
	inst := &Instruction{
		Code: code,
		Idx:  idx,
		Line: idx + 1,
	}

	if err := inst.parse(); nil != err {
		return nil, err
	}

	return inst, nil
}

// parse
func (inst *Instruction) parse() error {
	var err error

	// instruction type
	inst.Type, err = inst.parseType()
	if nil != err {
		return errors.Wrap(err, "failure identifying instruction type")
	}

	// parse instruction
	switch inst.Type {
	default:
		err = errors.Wrap(errUnknownInstruction, "instruction '%s' at line %d", inst.Type, inst.Line)
	case T_CONST:
		err = inst.parseConst()
	case T_LABEL:
		err = inst.parseLabel()
	case T_INSTR:
		err = inst.parseInstruction()
	case T_SUB:
		err = inst.parseSubroutine()
	case T_SUBEND:
		err = inst.endSubroutine()
	}
	if nil != err {
		return errors.Wrap(err, "failure parsing instruction type '%s' at line %d", inst.Code, inst.Line)
	}

	// generate opcodes
	err = inst.generateOpcodes()
	if nil != err {
		return errors.Wrap(err, "failure while generating opcodes for instruction '%s' at line %d", inst.Code, inst.Line)
	}

	return nil
}

func (inst *Instruction) generateOpcodes() error {

	return nil
}

// parseConst
func (inst *Instruction) parseConst() error {
	var err error

	parts := strings.Split(inst.Code, " ")
	if 2 != len(parts) {
		return errors.Wrap(errSyntaxError, "constant error '%s' at line %d", inst.Code, inst.Line)
	}

	inst.Label = parts[0]
	inst.Data, err = parseData(parts[1])
	if nil != err {
		return errors.Wrap(err, "error parsing constant data '%s' at line %d", inst.Code, inst.Line)
	}

	return nil
}

// parseLabel
func (inst *Instruction) parseLabel() error {
	parts := strings.Split(inst.Code, " ")
	if 1 != len(parts) {
		return errors.Wrap(errSyntaxError, "invalid label format '%s' at line %d", inst.Code, inst.Line)
	}
	inst.Label = parts[0]
	inst.Data = 0
	return nil
}

// parseInstruction
// Instructions consist of ` [instruction word] [instruction data]`
// Note, instructions begin with a non-zero ammount of whitespace.
// Not all instructions accept data:
//	* JMP: accepts a label name indexing a position in the program to jump to
//	* RUN: accepts a subroutine name indexing subroutine instructions
//	* LDAV, LDXV, LDYV, ADDV, SUBV, OUTV: accepts a $const data byte
//
//		* 42
//		* 0x66
//		* 0b00101010
func (inst *Instruction) parseInstruction() error {
	parts := strings.Split(strings.Trim(inst.Code, " "), " ")
	if 0 == len(parts) || 2 < len(parts) {
		return errors.Wrap(errSyntaxError, "invalid instruction format '%s' at line %d", inst.Code, inst.Line)
	}
	inst.Label = parts[0]

	inst.Data = 0
	if 2 == len(parts) {
		switch inst.Label {
		default:
			if data, ok := datMap[parts[1]]; ok {
				inst.Data = data
			} else if byt, err := parseData(parts[1]); nil != err {
				return errors.Wrap(err, "unknown data reference '%s' at line %d", parts[1], inst.Line)
			} else {
				inst.Data = byt
			}

		case "RUN":
			if idx, ok := subMap[parts[1]]; ok {
				inst.Data = idx
			} else {
				return errors.Wrap(errCompileError, "unknown subroutine reference '%s' at line %d", parts[1], inst.Line)
			}

		case "JMP":
			if idx, ok := jmpMap[parts[1]]; ok {
				inst.Data = idx
			} else {
				return errors.Wrap(errCompileError, "unknown label reference '%s' at line %d", parts[1], inst.Line)
			}

		case "JMPV":
			if data, ok := datMap[parts[1]]; ok {
				inst.Data = data
			} else if byt, err := parseData(parts[1]); nil != err {
				return errors.Wrap(err, "unknown data reference '%s' at line %d", parts[1], inst.Line)
			} else {
				inst.Data = byt
			}
		}
	}

	return nil
}

// parseSubroutine
func (inst *Instruction) parseSubroutine() error {
	name := strings.TrimRight(inst.Code, "{")
	if 1 != len(strings.Split(name, " ")) {
		return errors.Wrap(errSyntaxError, "invalid subroutine format '%s' at line %d", inst.Code, inst.Line)
	}
	inst.Label = name

	return nil
}

// endSubroutine
func (inst *Instruction) endSubroutine() error {
	parts := strings.Split(inst.Code, " ")
	if "}" != parts[0] {
		return errors.Wrap(errSyntaxError, "invalid subroutine return statement '%s' at line %d", inst.Code, inst.Line)
	}
	inst.Label = parts[0]
	inst.Data = 0

	return nil
}

// parseType
func (inst *Instruction) parseType() (instructionType, error) {
	var typ instructionType = T_UNK

	switch true {
	// Unknown instruction syntax, fail
	default:
		return typ, errors.Wrap(errParseError, "unexpected instruction '%s' at line %d", inst.Code, inst.Line)

	// Instructions
	case strings.HasPrefix(inst.Code, " "):
		typ = T_INSTR

	// Data
	case strings.HasPrefix(inst.Code, "$"):
		typ = T_CONST

	// Subroutine
	case strings.HasSuffix(inst.Code, "{"):
		typ = T_SUB

	// Subroutine-end
	case "}" == inst.Code:
		typ = T_SUBEND

	// Jump label
	case !strings.ContainsAny(inst.Code, " ") && "" != inst.Code:
		typ = T_LABEL
	}

	return typ, nil
}

// parseData
func parseData(dataStr string) (int, error) {
	var err error
	var data int64

	switch true {
	// binary
	case strings.HasPrefix(dataStr, "0b"):
		data, err = strconv.ParseInt(string(dataStr[2:]), 2, 8)
	// hex
	case strings.HasPrefix(dataStr, "0x"):
		data, err = strconv.ParseInt(string(dataStr[2:]), 16, 8)
	// decimal
	default:
		data, err = strconv.ParseInt(dataStr, 10, 8)
	}
	if nil != err {
		return 0, errors.Wrap(err, "failure parsing data '%s'", dataStr)
	}

	return int(data), nil
}
