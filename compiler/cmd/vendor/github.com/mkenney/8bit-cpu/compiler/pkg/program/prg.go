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

// Parse
func (prg *Program) parse() error {
	// Read source file.
	bytes, err := ioutil.ReadFile(prg.sourceFile)
	if nil != err {
		return errors.Wrap(err, "could not read source file '%s'", prg.sourceFile)
	}

	// Split source file into lines
	prg.Lines = strings.Split(string(bytes), "\n")

	// "lex" code
	for idx, line := range prg.Lines {
		// Strip comments.
		p := strings.Split(line, "#")
		code := strings.TrimRight(p[0], " \t")

		// All whitespace must a single tab.
		code = strings.ReplaceAll(code, "\t", " ")
		for strings.Contains(code, "  ") {
			code = strings.ReplaceAll(code, "  ", " ")
		}
		code = strings.ReplaceAll(code, " {", "{")

		// Parse instructions.
		if "" != strings.Trim(code, " \t") {
			prg.Code[idx] = code
			prg.Instructions[idx], err = newInstruction(code)
			if nil != err {
				return errors.Wrap(err, "failed to parse instruction '%s' at line %d", code, idx+1)
			}

			//
			prg.Instructions[idx].Idx = idx
			prg.Instructions[idx].Line = idx + 1
			prg.Instructions[idx].Inst = len(prg.Instructions) - 1

			// Update token indexes.
			switch prg.Instructions[idx].Type {
			default:
				return errors.Wrap(errUnknownInstruction, "instruction type '%s'", prg.Instructions[idx].Type)

			case T_CONST:
				// Data stored on ROM.
				datMap[prg.Instructions[idx].Code] = prg.Instructions[idx].Data

			case T_LABEL:
				// Label for JMP instructions.
				jmpMap[prg.Instructions[idx].Label] = idx

			case T_INSTR:
				// Instruction.

			case T_SUB:
				// Subroutine.
				prg.inSub = true
				prg.curSub = prg.Instructions[idx].Label
				subMap[prg.Instructions[idx].Label] = idx

			case T_SUBEND:
				// Subroutine end.
				prg.inSub = false
				prg.Instructions[idx].Label = prg.curSub
				prg.curSub = ""
			}
		}
	}
	return nil
}

func (prg *Program) Compile(dest string) error {
	var err error

	outf, err := os.Create(dest)
	if nil != err {
		return errors.Wrap(err, "could not create data file '%s'", dest)
	}
	defer outf.Close()

	bytes := []byte{}
	for k := range prg.Lines {
		if inst, ok := prg.Instructions[k]; ok {
			var instByte byte
			switch inst.Type {
			default:
				return errors.Wrap(errUnknownInstruction, "instruction type '%s'", inst.Type)

			case T_CONST:
				// Data stored on ROM.
				bytes = append(bytes, 0, 0)

			case T_LABEL:
				// Label for JMP instructions.
				instByte, err = opCode("LABEL")
				bytes = append(bytes, instByte, byte(inst.Data))

			case T_INSTR:
				// Instruction.
				instByte, err = opCode(inst.Label)
				if nil != err {
					log.WithError(err).Debug(err)
				}
				bytes = append(bytes, instByte, byte(inst.Data))
				log.WithFields(log.Fields{
					"inst":   inst.Label,
					"data":   inst.Data,
					"opcode": instByte,
				}).Debug("instruction data byte")

			case T_SUB:
				// Subroutine.
				bytes = append(bytes, 0, 0)

			case T_SUBEND:
				// Subroutine end.
				bytes = append(bytes, 0, 0)
			}
			if nil != err {
				return err
			}

			log.Debugf("compiled instruction %d: % -10s %02x:%02x\n", inst.Inst, `"`+strings.Trim(inst.Code, " ")+`"`, instByte, inst.Data)
		}
	}

	binSize := Kbit32 - len(bytes)
	for a := 0; a < binSize; a++ {
		bytes = append(bytes, 0)
	}

	err = binary.Write(outf, binary.BigEndian, bytes)
	if nil != err {
		return errors.Wrap(err, "could not write binary data to '%s'", dest)
	}

	return nil
}
