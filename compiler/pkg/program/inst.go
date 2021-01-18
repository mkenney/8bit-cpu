package program

import (
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

type Instruction struct {
	Inst  int
	Data  int
	Name  string
	Value int

	Line    string // Corresponding line from source file
	LineNum int    // Line number of source file
	Step    int    // Instruction set step number
	Type    instructionType
}

func newInstruction(ln int, line string) (*Instruction, error) {
	inst := &Instruction{
		Line:    line,
		LineNum: ln + 1,
	}

	err := inst.parse()
	if nil != err {
		return nil, errors.Wrap(err, "failure initializing instruction")
	}

	return inst, nil
}

func (inst *Instruction) parse() error {
	var err error

	err = inst.parseType()
	if nil != err {
		return errors.Wrap(err, "failure parsing instruction")
	}

	switch inst.Type {
	case TYPE_CONST:
		err = inst.parseConst()
	case TYPE_LABEL:
		err = inst.parseLabel()
	case TYPE_INSTR:
		err = inst.parseInstruction()
	case TYPE_SUB:
		err = inst.parseSubroutine()
	case TYPE_SUBEND:
		err = inst.endSubroutine()
	default:
		err = errors.Errorf("unknown instruction type '%d'", inst.Type)
	}
	if nil != err {
		return errors.Wrap(err, "failure parsing instruction type")
	}

	return nil
}

func (inst *Instruction) parseLabel() (*Instruction, error) {
	lParts := strings.Split(inst.Line, "\t")
	if 1 != len(lParts) {
		return nil, errors.Errorf("invalid label format: %s", inst.Line)
	}
	inst.Name = lParts[0]

	jmpMap[lParts[0]] = int(byt << inst.Step)
	stepMap[int(byt<<inst.Step)] = inst

	return nil, nil
}

func (inst *Instruction) parseType() error {
	switch true {

	// Instructions
	case strings.HasPrefix(inst.Name, "\t"):
		inst.Type = TYPE_INSTR

	// Data
	case strings.HasPrefix(inst.Name, "$"):
		inst.Type = TYPE_CONST

	// Subroutine
	case strings.HasSuffix(inst.Name, "{"):
		inst.Type = TYPE_SUB

	// Subroutine-end
	case "}" == inst.Name:
		inst.Type = TYPE_SUBEND

	// Jump label
	case !strings.ContainsAny(inst.Name, "\t") && "" != inst.Name:
		inst.Type = TYPE_LABEL

	// Unknown instruction syntax, fail
	default:
		return errors.Errorf("unexpected instruction '%s'", inst.Name)
	}

	return nil
}

func (inst *Instruction) parseConst() error {
	var err error

	lParts := strings.Split(inst.Line, "\t")
	if 2 != len(lParts) {
		return errors.Errorf("invalid data format: %s", inst.Line)
	}
	datMap[lParts[0]] = int(byt << inst.Step)

	var data int64
	if strings.HasPrefix(lParts[1], "0b") {
		data, err = strconv.ParseInt(string(lParts[1][2:]), 2, 8)
	} else if strings.HasPrefix(lParts[1], "0x") {
		data, err = strconv.ParseInt(string(lParts[1][2:]), 16, 8)
	} else {
		data, err = strconv.ParseInt(lParts[1], 10, 8)
	}
	if nil != err {
		log.WithError(err).Error(err)
	}
	inst.Data = int(data)

	return nil, nil
}

func (prg *Program) parseLabel(ln int) (*Instruction, error) {
	line := prg.lines[ln]
	lParts := strings.Split(line, "\t")
	if 1 != len(lParts) {
		return nil, errors.Errorf("invalid label format: %s", line)
	}

	inst := newInstruction(ln, line)
	jmpMap[lParts[0]] = int(byt << inst.Step)
	stepMap[int(byt<<inst.Step)] = inst

	return nil, nil
}

func (prg *Program) parseSubroutine(ln int) (*Instruction, error) {
	return nil, nil
}
func (prg *Program) endSubroutine(ln int) (*Instruction, error) {
	return nil, nil
}
