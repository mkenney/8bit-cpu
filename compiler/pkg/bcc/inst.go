package bcc

import (
	"strings"

	"github.com/bdlm/errors/v2"
)

// instruction represents a line of code and manages tokenization and
// compilation of that line.
type instruction struct {
	ln     int
	line   string
	tokens []*tok
	op     *oper
}

func (inst *instruction) Type() tokenType {
	return inst.tokens[0].typ
}

func newInst(ln int, line string) (*instruction, error) {
	inst := &instruction{
		ln:   ln,
		line: line,
	}

	err := inst.tokenize()
	if nil != err {
		return nil, errors.Wrap(err, "error tokenizing instruction")
	}

	op, err := newOp(inst.tokens)
	if nil != err {
		return nil, errors.Wrap(err, "error generating instruction opcodes")
	}
	inst.op = op

	return inst, nil
}

func (inst *instruction) Line() string {
	return inst.line
}

func (inst *instruction) compile() ([]byte, error) {
	return inst.op.byts, nil
}

func (inst *instruction) tokenize() error {
	tokens := []*tok{}
	parts := strings.Split(inst.line, " ")
	for pos, prt := range parts {
		if "" == prt {
			continue
		}

		tkn, err := newToken(inst.ln, pos, prt)
		if nil != err {
			return errors.Wrap(err, "token failure")
		}

		tokens = append(tokens, tkn)
	}
	inst.tokens = tokens

	return nil
}

//type Instruction struct {
//	Code  string // Line from source file
//	Data  int    // Data portion of the instruction code
//	Idx   int    // Source file index
//	Line  int    // Line number from source file
//	Label string
//
//	Inst    int    // Instruction address in compiled source
//	Opcodes []byte // Opcodes in compiled source
//
//	Step int // Instruction set step number
//	Type int
//}
