// package main defines the executable for the bcc (bit code compiler) compiler.
package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	prgm "github.com/mkenney/8bit-cpu/compiler/pkg/program"

	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

type instructionType int

const (
	// Data stored on ROM
	TYP_DATA instructionType = iota

	// Label for JMP instructions.
	TYP_LABEL

	// Instruction
	TYP_INSTR

	// Subroutine
	TYP_SUB

	// Subroutine end
	TYP_SUBEND
)

type program []instruction

type instruction struct {
	Inst int
	Data int
	Name string

	Step int
	Typ  instructionType
}

var prg program

type tokenMap map[string]int

var jmpMap = map[string]int{}
var subMap = map[string]int{}
var stepMap = map[int]string{}
var varMap = map[string]int{}
var subStack = []int{}
var callStack = []int{}

var byt int = 0

func main() {
	var err error

	prg, err := prgm.New(os.Args[1])
	if nil != err {
		log.WithError(err).Fatal("failed to initialize program parser")
	}

	err = prg.parse()
	if nil != err {
		return nil, errors.Wrap(err, "failed to parse source file '%s'", os.Args[1])
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if nil != err {
		log.WithError(err).Fatal(err)
	}
	file := string(bytes)

	file = strings.ReplaceAll(file, " ", "\t")
	for strings.Contains(file, "\t\t") {
		file = strings.ReplaceAll(file, "\t\t", "\t")
	}

	l := strings.Split(file, "\n")
	lines := []string{}
	for _, v := range l {
		p := strings.Split(v, "#")
		if "" != p[0] {
			lines = append(lines, p[0])
			// DEBUG
			fmt.Println(p[0])
		}
	}

	inSubroutine := false
	for _, line := range lines {
		// strip comments
		lparts := strings.Split(line, "#")
		line = lparts[0]

		// all whitespace must be tabs
		line = strings.ReplaceAll(line, " ", "\t")
		for strings.Contains(line, "\t\t") {
			line = strings.ReplaceAll(line, "\t\t", "\t")
		}

		// ignore blank lines
		if "" == strings.Trim(line, " \t") {
			continue
		}

		// tokenize current line
		lparts = strings.Split(line, "\t")
		tok := instruction{
			Step: len(prg),
		}

		// labels
		if "" != string(lparts[0]) {
			tok.Name = lparts[0]
			tok.Inst = instmap[tok.Name]

			switch true {

			// Constants
			case strings.HasPrefix(line, "$"):
				tok.Typ = TYP_DATA
				varMap[lparts[0]] = byt << tok.Step
				if len(lparts) > 1 {
					var d int64
					if strings.HasPrefix(lparts[1], "0b") {
						d, err = strconv.ParseInt(string(lparts[1][2:]), 2, 8)
					} else if strings.HasPrefix(lparts[1], "0x") {
						d, err = strconv.ParseInt(string(lparts[1][2:]), 16, 8)
					} else {
						d, err = strconv.ParseInt(lparts[1], 10, 8)
					}
					if nil != err {
						log.WithError(err).Error(err)
					}
					tok.Data = int(d)
				}

			// Subroutine start
			case strings.HasSuffix(line, "{"):
				tok.Typ = TYP_SUB
				if inSubroutine {
					log.WithField("sub", lparts[0]).Fatal("nested subroutine call")
				}
				inSubroutine = true

				log.WithField("sub", lparts[0]).Info("Begin subroutine")

				subMap[lparts[0]] = byt << tok.Step
				stepMap[byt<<tok.Step] = lparts[0]
				subStack = append(subStack, byt<<tok.Step)

				tok.Typ = TYP_SUB

			// Subroutine end
			case "}" == line:
				tok.Typ = TYP_SUBEND
				if !inSubroutine {
					log.WithField("token", line).Fatal("unexpected token")
				}
				inSubroutine = false

				endSub := subStack[len(subStack)-1]
				subStack = subStack[:len(subStack)-1]
				log.WithField("sub", stepMap[endSub]).Info("End subroutine")

			// JMP label
			default:
				tok.Typ = TYP_LABEL
				jmpMap[tok.Name] = byt << tok.Step
				if len(lparts) >= 2 {
					d, err := strconv.ParseInt(lparts[1], 10, 64)
					if nil != err {
						log.WithError(err).Error(err)
					}
					tok.Data = int(d)
				}
				tok.Typ = TYP_LABEL
			}

		} else {
			tok.Typ = TYP_INSTR
			tok.Name = lparts[1]
			tok.Inst = instmap[tok.Name]
			if len(lparts) >= 3 {
				if "RUN" == tok.Name {
					tok.Data = subMap[lparts[2]]
				} else if "JMP" == tok.Name {
					tok.Data = jmpMap[lparts[2]]
				} else {
					d, err := strconv.ParseInt(lparts[2], 10, 64)
					if nil != err {
						log.WithError(err).Error(err)
					}
					tok.Data = int(d)
				}
			}
		}

		prg = append(prg, tok)
	}

	log.WithField("prg", prg).Info("tokenized")

	compile(prg)
}

func compile(prg []instruction) {
	outf, err := os.Create("bcc.bin")
	if nil != err {
		log.WithError(err).Fatal(err)
	}
	defer outf.Close()

	bytes := []byte{}
	for _, tok := range prg {
		bytes = append(bytes, byte(tok.Inst), byte(tok.Data))
		fmt.Printf("%d:%d %s\n", byte(tok.Inst), byte(tok.Data), tok.Name)
	}

	err = binary.Write(outf, binary.BigEndian, bytes)
	if nil != err {
		log.WithError(err).Fatal(err)
	}

	cnt := 0
	for _, b := range bytes {
		if cnt > 15 {
			cnt = 0
			fmt.Printf("\n")
		}
		if cnt == 8 {
			fmt.Printf(": ")
		}
		fmt.Printf("%02x ", b)
		cnt++
	}
}

var bitmap = map[int]string{
	0:  "RUN", // [LABEL]
	1:  "HLT",
	2:  "RST",
	3:  "NOP",
	4:  "SLOP",
	5:  "LDA",   // Load BYTE into register A // [BYTE]
	6:  "LDX",   // Load BYTE into register X // [BYTE]
	7:  "LDY",   // Load BYTE into register Y // [BYTE]
	8:  "LDA,G", // Load register RID into register A // [RID]
	9:  "LDX,G", // Load register RID into register X // [RID]
	10: "LDY,G", // Load register RID into register Y // [RID]
	11: "LDA,M", // Load ROM address ADDR into register A // [ADDR]
	12: "LDX,M", // Load ROM address ADDR into register X // [ADDR]
	13: "LDY,M", // Load ROM address ADDR into register Y // [ADDR]
	14: "LDA,R", // Load RAM address ADDR into register A // [ADDR]
	15: "LDX,R", // Load RAM address ADDR into register X // [ADDR]
	16: "LDY,R", // Load RAM address ADDR into register Y // [ADDR]
	17: "OUT",   // [BYTE]
	18: "OUT,G", // [RID]
	19: "OUT,M", // [ADDR]
	20: "OUT,R", // [ADDR]
	21: "PSH",   // [BYTE]
	22: "PSH,G", // [RID]
	23: "PSH,M", // [ADDR]
	24: "PSH,R", // [ADDR]
	25: "PUL",
	26: "PUL,G",     // [RID]
	27: "PUL,R",     // [ADDR]
	28: "ST,[ADDR]", // [DATA]
	29: "STG,[RID]", // [ADDR]
	30: "STM",       // [ADDR]
	31: "STR",       // [ADDR]
	32: "ADD",       // [BYTE]
	33: "ADDG",      // [RID]
	34: "ADDM",      // [ADDR]
	35: "ADDR",      // [ADDR]
	36: "SUB",       // [BYTE]
	37: "SUBG",      // [RID]
	38: "SUBM",      // [ADDR]
	39: "SUBR",      // [ADDR]
	40: "JMP",       // [LABEL]
	41: "JMPG",      // [RID]
	42: "JMPM",      // [ADDR]
	43: "JMPR",      // [ADDR]
}

var instmap = map[string]int{
	"RUN":       0, //  [LABEL]
	"HLT":       1,
	"RST":       2,
	"NOP":       3,
	"SLOP":      4,
	"LDA":       5,  // Load BYTE into register A //  [BYTE]
	"LDX":       6,  // Load BYTE into register X //  [BYTE]
	"LDY":       7,  // Load BYTE into register Y //  [BYTE]
	"LDA,G":     8,  // Load register RID into register A //  [RID]
	"LDX,G":     9,  // Load register RID into register X //  [RID]
	"LDY,G":     10, // Load register RID into register Y //  [RID]
	"LDA,M":     11, // Load ROM address ADDR into register A //  [ADDR]
	"LDX,M":     12, // Load ROM address ADDR into register X //  [ADDR]
	"LDY,M":     13, // Load ROM address ADDR into register Y //  [ADDR]
	"LDA,R":     14, // Load RAM address ADDR into register A //  [ADDR]
	"LDX,R":     15, // Load RAM address ADDR into register X //  [ADDR]
	"LDY,R":     16, // Load RAM address ADDR into register Y //  [ADDR]
	"OUT":       17, //  [BYTE]
	"OUTG":      18, //  [RID]
	"OUTM":      19, //  [BYTE]
	"OUTR":      20, //  [BYTE]
	"PSH":       21, //  [BYTE]
	"PSHG":      22, //  [RID]
	"PSHM":      23, //  [BYTE]
	"PSHR":      24, //  [BYTE]
	"PUL":       25,
	"PULG":      26, //  [RID]
	"PULR":      27, //  [BYTE]
	"ST,[ADDR]": 28, //  [DATA]
	"STG,[RID]": 29, //  [ADDR]
	"STM":       30, //  [ADDR]
	"STR":       31, //  [ADDR]
	"ADD":       32, //  [BYTE]
	"ADDG":      33, //  [RID]
	"ADDM":      34, //  [ADDR]
	"ADDR":      35, //  [ADDR]
	"SUB":       36, //  [BYTE]
	"SUBG":      37, //  [RID]
	"SUBM":      38, //  [ADDR]
	"SUBR":      39, //  [ADDR]
	"JMP":       40, //  [LABEL]
	"JMPG":      41, //  [RID]
	"JMPM":      42, //  [ADDR]
	"JMPR":      43, //  [ADDR]
}
