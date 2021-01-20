package program

import (
	"encoding/binary"
	"fmt"
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

	datMap map[string]int // data map of $const => index
	jmpMap map[string]int // label map of label => index
	subMap map[string]int // subroutine map of subroutine{ => index
}

// New
func New(sourceFile string) (*Program, error) {
	prg := &Program{
		Lines:        []string{},
		Code:         map[int]string{},
		Instructions: map[int]*Instruction{},
		sourceFile:   sourceFile,

		datMap: map[string]int{},
		jmpMap: map[string]int{},
		subMap: map[string]int{},
	}

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
		log.WithError(err).Error(err)
	}

	// Split source file into lines
	prg.Lines = strings.Split(string(bytes), "\n")

	// "lex" code
	for idx, line := range prg.Lines {
		// Strip comments.
		p := strings.Split(line, "#")
		code := strings.TrimRight(p[0], " \t")

		// All whitespace must a single space.
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
			prg.Instructions[idx].Line = idx + 1
			prg.Instructions[idx].Inst = len(prg.Instructions) - 1

			// Update token indexes.
			switch prg.Instructions[idx].Type {
			default:
				return errors.Wrap(errUnknownInstruction, "instruction type '%s'", prg.Instructions[idx].Type)

			case T_CONST:
				// Data stored on ROM.
				prg.datMap[prg.Instructions[idx].Code] = idx

			case T_LABEL:
				// Label for JMP instructions.
				prg.jmpMap[prg.Instructions[idx].Code] = idx

			case T_INSTR:
				// Instruction.

			case T_SUB:
				// Subroutine.
				prg.inSub = true
				prg.curSub = prg.Instructions[idx].Label
				prg.subMap[prg.Instructions[idx].Label] = idx

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

func (prg *Program) Compile() error {
	fmt.Printf("\n\n\n")

	outf, err := os.Create("bcc.bin")
	if nil != err {
		log.WithError(err).Fatal(err)
	}
	defer outf.Close()

	bytes := []byte{}
	for k := range prg.Lines {
		if inst, ok := prg.Instructions[k]; ok {
			bytes = append(bytes, byte(inst.Inst), byte(inst.Data))
		}
	}

	binSize := Kbit32 - len(bytes)
	for a := 0; a < binSize; a++ {
		bytes = append(bytes, 0)
	}

	err = binary.Write(outf, binary.BigEndian, bytes)
	if nil != err {
		log.WithError(err).Fatal(err)
	}

	cnt := 0
	for a, b := range bytes {
		if a > len(prg.Instructions)*2 {
			break
		}
		if 0 == a%15 {
			cnt = 0
			fmt.Printf("\n% 8x", a)
		}
		if 0 == cnt%8 {
			fmt.Printf(":  ")
		}
		fmt.Printf("%02x ", b)
		cnt++
	}
	fmt.Printf("\n\n\n")

	return nil
}
