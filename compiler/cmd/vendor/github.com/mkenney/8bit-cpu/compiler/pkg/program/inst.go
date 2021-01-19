package program

import (
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

// Instruction
type Instruction struct {
	Code  string // Line from source file
	Data  int
	Idx   int // Source file index
	Line  int // Line number from source file
	Label string

	Inst    int    // Instruction address in compiled source
	Opcodes []byte // Opcodes in compiled source

	Step int // Instruction set step number
	Type instructionType
}

// newInstruction
func newInstruction(code string) (*Instruction, error) {
	inst := &Instruction{Code: code}

	if err := inst.parse(); nil != err {
		return nil, err
	}

	return inst, nil
}

// parse
func (inst *Instruction) parse() error {
	var err error

	// instruction type
	err = inst.parseType()
	if nil != err {
		return errors.Wrap(err, "failure identifying instruction type")
	}

	// parse instruction
	switch inst.Type {
	default:
		err = errors.Wrap(errUnknownInstruction, "instruction '%s'", inst.Type)
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
		return errors.Wrap(err, "failure parsing instruction type '%s'", inst.Code)
	}

	// generate opcodes
	err = inst.generateOpcodes()
	if nil != err {
		return errors.Wrap(err, "failure while generating opcodes for instruction '%s'", inst.Code)
	}

	return nil
}

func (inst *Instruction) generateOpcodes() error {

	return nil
}

// parseConst
func (inst *Instruction) parseConst() error {
	var err error

	parts := strings.Split(inst.Code, "\t")
	if 2 != len(parts) {
		return errors.Wrap(errSyntaxError, "constant error '%s'", inst.Code)
	}

	inst.Data, err = parseData(parts[1])
	if nil != err {
		return errors.Wrap(err, "error parsing constant data '%s'", inst.Code)
	}

	return nil
}

// parseLabel
func (inst *Instruction) parseLabel() error {
	parts := strings.Split(inst.Code, "\t")
	if 1 != len(parts) {
		return errors.Wrap(errSyntaxError, "invalid label format '%s'", inst.Code)
	}
	inst.Label = parts[0]
	inst.Data = 0
	return nil
}

// parseInstruction
func (inst *Instruction) parseInstruction() error {
	var err error
	parts := strings.Split(strings.Trim(inst.Code, " \t"), "\t")
	if 1 > len(parts) {
		return errors.Wrap(errSyntaxError, "invalid instruction format '%s'", inst.Code)
	}
	inst.Label = parts[0]
	if 2 == len(parts) {
		inst.Data, err = parseData(parts[1])
		if nil != err {
			log.WithError(errors.Wrap(err, "error parsing instruction data '%s'", inst.Code)).Warn("non-numeric instruction data")
		}
	}

	return nil
}

// parseSubroutine
func (inst *Instruction) parseSubroutine() error {
	name := strings.TrimRight(inst.Code, "{")
	if 1 != len(strings.Split(name, "\t")) {
		return errors.Wrap(errSyntaxError, "invalid subroutine format '%s'", inst.Code)
	}
	inst.Label = name

	return nil
}

// endSubroutine
func (inst *Instruction) endSubroutine() error {
	parts := strings.Split(inst.Code, "\t")
	if "}" != parts[0] {
		return errors.Wrap(errSyntaxError, "invalid subroutine return statement '%s'", inst.Code)
	}
	inst.Label = parts[0]

	return nil
}

// parseType
func (inst *Instruction) parseType() error {
	switch true {
	// Unknown instruction syntax, fail
	default:
		return errors.Wrap(errParseError, "unexpected instruction '%s'", inst.Code)

	// Instructions
	case strings.HasPrefix(inst.Code, "\t"):
		inst.Type = T_INSTR

	// Data
	case strings.HasPrefix(inst.Code, "$"):
		inst.Type = T_CONST

	// Subroutine
	case strings.HasSuffix(inst.Code, "{"):
		inst.Type = T_SUB

	// Subroutine-end
	case "}" == inst.Code:
		inst.Type = T_SUBEND

	// Jump label
	case !strings.ContainsAny(inst.Code, "\t") && "" != inst.Code:
		inst.Type = T_LABEL
	}

	return nil
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
