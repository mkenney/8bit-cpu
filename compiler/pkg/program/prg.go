package program

import (
	"io/ioutil"
	"strings"

	"github.com/bdlm/errors/v2"
	"github.com/bdlm/log/v2"
)

type Program struct {
	instructions []*Instruction
	inSubroutine bool
	lines        []string
	sourceFile   string

	datMap  map[string]int       // data map
	jmpMap  map[string]int       // labels
	stepMap map[int]*Instruction // step -> inst map
	subMap  map[string]int       // subroutine map
}

func New(sourceFile string) (*Program, error) {
	prg := &Program{
		instructions: []*Instruction{},
		sourceFile:   sourceFile,
	}

	return prg, nil
}

func (prg *Program) Parse() error {
	// Read sourceFile.
	bytes, err := ioutil.ReadFile(prg.sourceFile)
	if nil != err {
		log.WithError(err).Error(err)
	}

	// All whitespace must be tabs.
	file := strings.ReplaceAll(string(bytes), " ", "\t")
	for strings.Contains(file, "\t\t") {
		file = strings.ReplaceAll(file, "\t\t", "\t")
	}

	l := strings.Split(file, "\n")
	lines := []string{}
	for _, line := range l {
		// Strip comments.
		p := strings.Split(line, "#")
		lines = append(lines, p[0])
	}

	return prg.parse(lines)
}

func (prg *Program) parse(lines []string) error {
	var err error
	prg.inSubroutine = false

	for ln, line := range lines {
		// Ignore blank lines.
		if "" == strings.Trim(line, "\t") {
			continue
		}

		// Parse current instruction.
		inst, err := newInstruction(ln, line)
		if nil != err {
			return errors.Wrap(err, "error parsing program instruction")
		}

		inst.Step = len(prg.instructions)
		prg.instructions = append(prg.instructions, inst)
	}

	return nil
}
