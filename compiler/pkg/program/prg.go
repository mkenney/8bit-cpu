package program

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

type Program struct {
	Lines        []string
	Code         map[int]string
	Instructions map[int]*Instruction

	sourceFile string //
	inSub      bool   //
	curSub     string //

}

var datMap map[string]int // data map of $const => index
var jmpMap map[string]int // label map of label => index
var subMap map[string]int // subroutine map of subroutine{ => index

// New
func New(sourceFile string) (*Program, error) {
	prg := &Program{
		Lines:        []string{},
		Code:         map[int]string{},
		Instructions: map[int]*Instruction{},
		sourceFile:   sourceFile,
	}

	datMap = map[string]int{}
	jmpMap = map[string]int{}
	subMap = map[string]int{}

	if err := prg.parse(); nil != err {
		return nil, errors.Wrap(err, "parse error")
	}

	return prg, nil
}

// Parse parses the source file, performing "lexical analysis"... just a bunch
// of strings.Split and if statements :)
func (prg *Program) parse() error {
	// Read source file and split into lines.
	bytes, err := ioutil.ReadFile(prg.sourceFile)
	if nil != err {
		return errors.Wrap(err, "could not read source file '%s'", prg.sourceFile)
	}
	prg.Lines = strings.Split(string(bytes), "\n")

	// Inspect each line, tokenizing instructions into program.Instruction{}
	// structs.
	for idx, line := range prg.Lines {
		// Strip comments.
		p := strings.Split(line, "#")
		code := strings.TrimRight(p[0], " \t")

		// All whitespace must be a single space.
		code = strings.ReplaceAll(code, "\t", " ")
		for strings.Contains(code, "  ") {
			code = strings.ReplaceAll(code, "  ", " ")
		}
		code = strings.ReplaceAll(code, " {", "{")

		// Remaining non-blank lines are code instructions.
		if "" != strings.Trim(code, " \t") {
			prg.Code[idx] = code
			prg.Instructions[idx], err = newInstruction(code, idx)
			if nil != err {
				return errors.Wrap(err, "failed to parse instruction '%s' at line %d", code, idx+1)
			}

			// Program counter index associated with this instruction.
			prg.Instructions[idx].Inst = len(prg.Instructions) - 1

			// Update global instruction indexes.
			switch prg.Instructions[idx].Type {
			default:
				return errors.Wrap(errUnknownInstruction, "instruction type '%s'", prg.Instructions[idx].Type)

			case T_CONST:
				// ROM image indexes for $labeled data.
				datMap[prg.Instructions[idx].Label] = prg.Instructions[idx].Data

			case T_LABEL:
				// indexes for JMP instructions.
				jmpMap[prg.Instructions[idx].Label] = idx

			case T_INSTR:
				// instruction statements.

			case T_SUB:
				// subroutine label indexes.
				prg.inSub = true
				prg.curSub = prg.Instructions[idx].Label
				subMap[prg.Instructions[idx].Label] = idx

			case T_SUBEND:
				// subroutine end label '}' injects code that uses the call stack
				// to return from a sub, does not need mapping.
				prg.Instructions[idx].Label = prg.curSub
				prg.inSub = false
				prg.curSub = ""
			}
		}
	}

	return nil
}

// Compile puts all the bits in their places and writes them to a file at `dest`.
func (prg *Program) Compile(dest string) error {
	var err error

	outf, err := os.Create(dest)
	if nil != err {
		return errors.Wrap(err, "could not create data file '%s'", dest)
	}
	defer outf.Close()

	bytes := []byte{}
	for idx := range prg.Lines {
		if inst, ok := prg.Instructions[idx]; ok {
			var byt byte
			switch inst.Type {
			default:
				return errors.Wrap(errUnknownInstruction, "instruction type '%s'", inst.Type)

			case T_CONST:
				// Data stored on ROM.
				// 0x00 [data value]
				bytes = append(bytes, 0, byte(inst.Data))

			case T_LABEL:
				// Label for JMP instructions.
				byt, err = opCode("NOP")
				if nil != err {
					return errors.Wrap(err, "OPCODE")
				}
				// [instruction] [data value]
				bytes = append(bytes, byt, 0)

			case T_INSTR:
				switch inst.Label {
				// Instruction.
				default:
					byt, err = opCode(inst.Label)
					if nil != err {
						return errors.Wrap(err, "OPCODE")
					}
					// [instruction] [data value]
					bytes = append(bytes, byt, byte(inst.Data))
					log.WithFields(log.Fields{
						"inst":   inst.Label,
						"data":   inst.Data,
						"opcode": byt,
					}).Debug("instruction data byte")

				// add current position to stack, jump to postion referenced by
				// instruction at index `inst.Data`
				case "RUN":
					if sub, ok := prg.Instructions[inst.Data]; ok {
						byt, err = opCode("JMPV")
						if nil != err {
							return errors.Wrap(err, "OPCODE")
						}
						bytes = append(bytes, byt, byte(sub.Idx))
					} else {
						return errors.Wrap(errUnknownSubroutine, "unknown subroutine '%s': %d", inst.Code, inst.Data)
					}

				case "JMPV":
				}

			case T_SUB:
				// Subroutine, implicit jump to subroutine label.
				//	* Push current program counter value onto stack
				//	* Jump to subroutine header index
				//
				// JMP [subroutine instruction postion]

				// push idx+1 into stack
				byt, err = opCode("PSHV")
				if nil != err {
					log.WithError(err).Debug(err)
					return errors.Wrap(err, "OPCODE")
				}
				bytes = append(bytes, byt, byte(inst.Idx+1))

				// jump to the subroutine index
				byt, err = opCode("JMP")
				if nil != err {
					return errors.Wrap(err, "OPCODE")
				}
				bytes = append(bytes, byt, byte(inst.Data))

			case T_SUBEND:
				// Subroutine end, implicit return from subroutine instruction block.
				//	* Pop last program value off of stack
				//	* Jump to program counter value
				//
				// JMP [program counter value]

				// load the last stack value into the program counter register
				byt, err = opCode("JMPS")
				if nil != err {
					log.WithError(err).Debug(err)
					return errors.Wrap(err, "OPCODE")
				}
				bytes = append(bytes, byt, byte(inst.Idx+1))
			}
			if nil != err {
				return err
			}

			log.Debugf("compiled instruction %d: % -10s %02x:%02x\n", inst.Inst, `"`+strings.Trim(inst.Code, " ")+`"`, byt, inst.Data)
		}
	}
	log.Infof("compiled code size is %d bytes", len(bytes))

	binSize := Kbit32 - len(bytes)
	log.Infof("padding ROM image to %d additional bytes (%d total)", binSize, Kbit32)
	for a := 0; a < binSize; a++ {
		bytes = append(bytes, 0)
	}

	err = binary.Write(outf, binary.BigEndian, bytes)
	if nil != err {
		return errors.Wrap(err, "could not write binary data to '%s'", dest)
	}

	return nil
}
