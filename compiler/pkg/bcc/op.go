package bcc

import (
	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

// init populates opMap.
func init() {
	for pcid, op := range opTable {
		op.pcid = byte(pcid)
		opMap[op.name] = op
	}
}

func newOp(tokens []*tok) (*oper, error) {
	var err error

	if 0 == len(tokens) {
		return nil, errors.Errorf("no tokens given")
	}

	op := &oper{
		tokens: tokens,
	}

	err = op.tokenize()
	if nil != err {
		return nil, errors.Wrap(err, "error generating opcodes")
	}

	return op, nil
}

func (op *oper) tokenize() error {
	for _, tkn := range op.tokens {
		switch tkn.pos {
		default:
			log.WithField("token", tkn).Debug("GOT HERE")
			return errors.Errorf("invalid token postion '%d'", tkn.pos)
		case 0:
			// switch tkn.typ {
			// case TOK_CONST:
			// case TOK_LABEL:
			// case TOK_SUB:
			// case TOK_SUBEND:
			// }
			op.byts = append(op.byts, byte(tkn.ln))
		case 1:
			switch tkn.typ {
			case TOK_LIT:
				op.byts = append(op.byts, tkn.dat)
			case TOK_OP:
				ref, ok := opMap[tkn.tkn]
				if !ok {
					return errors.Errorf("unknown operation '%s'", tkn.tkn)
				}
				op.byts = append(op.byts, ref.pcid)
			}
		case 2:
			switch tkn.typ {
			case TOK_LIT:
				op.byts = append(op.byts, tkn.dat)
			case TOK_CREF:
				byt, ok := constMap[tkn.tkn]
				if !ok {
					return errors.Errorf("unknown reference '%s'", tkn.tkn)
				}
				op.byts = append(op.byts, byt)
			case TOK_LREF:
				op.byts = append(op.byts, byte(tkn.ln))
			}
		}
	}

	return nil
}

type oper struct {
	byts   []byte
	tokens []*tok
	// Operation name as defined in `opTable`.
	name string
	// Whether this operation accepts param data.
	hasParam bool
	// Program counter Id.
	pcid byte
}

var opMap map[string]*oper = map[string]*oper{}

var opTable []*oper = []*oper{

	// internal
	&oper{name: string(TOK_NIL), hasParam: false},
	&oper{name: string(TOK_CONST), hasParam: false},
	&oper{name: string(TOK_CREF), hasParam: false},
	&oper{name: string(TOK_LIT), hasParam: false},
	&oper{name: string(TOK_LABEL), hasParam: false},
	&oper{name: string(TOK_LREF), hasParam: false},
	&oper{name: string(TOK_OP), hasParam: false},
	&oper{name: string(TOK_SUB), hasParam: false},
	&oper{name: string(TOK_SUBEND), hasParam: false},

	// system
	&oper{name: "HLT", hasParam: false},  // Halt system clock signal
	&oper{name: "RST", hasParam: false},  // Reset all system registers
	&oper{name: "NOP", hasParam: false},  // No-op, use 1 instruction cycle
	&oper{name: "SLOP", hasParam: false}, // Slow no-op, use 16 instruction cycles

	// math
	&oper{name: "ADDV", hasParam: true},  // Add a $const or literal to register A
	&oper{name: "ADDX", hasParam: false}, // Add register X to register A
	&oper{name: "ADDY", hasParam: false}, // Add register Y to register A

	&oper{name: "SUBV", hasParam: true},  // Subtract a $const or literal from register A
	&oper{name: "SUBX", hasParam: false}, // Subtract register X from register A
	&oper{name: "SUBY", hasParam: false}, // Subtract register Y from register A

	// branching logic
	&oper{name: "RUN", hasParam: true}, // Execute a subroutine. Is encoded as a JMPV operation

	&oper{name: "JMP", hasParam: true},   // `JMP [label]`  - Jump to [label]:       Load a label index into the program counter
	&oper{name: "JMPV", hasParam: true},  // `JMPV [value]` - Jump to [value]:       Load a $const or literal value into the program counter
	&oper{name: "JMPA", hasParam: false}, // `JMPA`         - Jump to register A:    Load register A into the program counter
	&oper{name: "JMPX", hasParam: false}, // `JMPX`         - Jump to register X:    Load register X into the program counter
	&oper{name: "JMPY", hasParam: false}, // `JMPY`         - Jump to register Y:    Load register Y into the program counter
	&oper{name: "JMPS", hasParam: false}, // `JMPS`         - Jump to stack pointer: Load the last stack value into the program counter

	// data
	&oper{name: "LDAV", hasParam: true},  // Load a $const or literal value into register A
	&oper{name: "LDAX", hasParam: false}, // Load register X into register A
	&oper{name: "LDAY", hasParam: false}, // Load register Y into register A

	&oper{name: "LDXV", hasParam: true},  // Load a $const or literal value into register X
	&oper{name: "LDXA", hasParam: false}, // Load register A into register X
	&oper{name: "LDXY", hasParam: false}, // Load register Y into register X

	&oper{name: "LDYV", hasParam: true},  // Load a $const or literal value into register Y
	&oper{name: "LDYA", hasParam: false}, // Load register A into register Y
	&oper{name: "LDYX", hasParam: false}, // Load register X into register Y

	// stack
	&oper{name: "PSHV", hasParam: true},  // Push a $const or literal value onto the stack
	&oper{name: "PSHA", hasParam: false}, // Push register A onto the stack
	&oper{name: "PSHX", hasParam: false}, // Push register X onto the stack
	&oper{name: "PSHY", hasParam: false}, // Push register Y onto the stack
	&oper{name: "PSHP", hasParam: false}, // Push the current program counter onto the stack

	&oper{name: "POPA", hasParam: false}, // Pop a stack value into register A
	&oper{name: "POPX", hasParam: false}, // Pop a stack value into register X
	&oper{name: "POPY", hasParam: false}, // Pop a stack value into register Y
	&oper{name: "POPP", hasParam: false}, // Pop the last value from the stack into the program counter

	// output
	&oper{name: "OUTV", hasParam: true},  // Send a value to the output register
	&oper{name: "OUTA", hasParam: false}, // Send the A register to the output register
	&oper{name: "OUTX", hasParam: false}, // Send the X register to the output register
	&oper{name: "OUTY", hasParam: false}, // Send the Y register to the output register
}
