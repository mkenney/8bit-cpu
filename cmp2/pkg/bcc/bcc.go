package bcc

import (
	"io/ioutil"
	"strings"

	"github.com/bdlm/errors"
)

const (
	// Kbit32
	Kbit32 = 32768
)

type Bcc interface {
	Parse() error
	Compile() error
	String() string
}

func New(sourceFile, destFile string) (Bcc, error) {
	return &bcc{
		sourceFile: sourceFile,
		destFile:   destFile,
	}, nil
}

var constMap map[string]byte // data map of $const => byte
var jmpMap map[string]int    // location map of label => program index
var subMap map[string]int    // subroutine map of subroutine{ => program index

type bcc struct {
	// System files
	sourceFile string
	destFile   string

	// Program
	prg [Kbit32]byte

	// Instruction module table
	inst []byte

	// maps and indexes
	lines        []string       // [idx]line from source
	instructions []*instruction //
}

func (bcc *bcc) compile() error {
	return nil
}

func (bcc *bcc) readSource() error {
	// Read source file and split into lines.
	bytes, err := ioutil.ReadFile(bcc.sourceFile)
	if nil != err {
		return errors.Wrap(err, "could not read source file '%s'", bcc.sourceFile)
	}
	bcc.lines = strings.Split(string(bytes), "\n")
	return nil
}

func (bcc *bcc) lex() error {
	// Inspect each line, tokenizing all elements.
	for idx, line := range bcc.lines {
		// Strip comments.
		p := strings.Split(line, "#")
		code := strings.TrimRight(p[0], " \t")

		// All whitespace must be a single space.
		code = strings.ReplaceAll(code, "\t", " ")
		for strings.Contains(code, "  ") {
			code = strings.ReplaceAll(code, "  ", " ")
		}
		code = strings.ReplaceAll(code, " {", "{")

		// Remaining non-blank lines are code instructions. Tokenize instructions,
		// populate maps.
		if "" != strings.Trim(code, " \t") {

			inst, err := newInst(idx+1, line)
			if nil != err {
				return errors.Wrap(err, "instruction tokenization failure on line %d: '%s'", idx+1, line)
			}

			// populate instruction maps
			switch inst.tokens[0].typ {
			case TOK_CONST:
				constMap[inst.tokens[0].tkn] = inst.tokens[0].dat
			case TOK_LABEL:
				jmpMap[inst.tokens[0].tkn] = len(bcc.instructions)
			case TOK_SUB:
				subMap[inst.tokens[0].tkn] = len(bcc.instructions)
			}

			bcc.instructions = append(bcc.instructions, inst)
		}
	}

	return nil
}

// parse parses the source file, performing "lexical analysis"... just a bunch
// of strings.Split and if statements :)
func (bcc *bcc) parse() error {
	err := bcc.readSource()
	if nil != err {
		return errors.Wrap(err, "error reading source file")
	}

	err = bcc.lex()
	if nil != err {
		return errors.Wrap(err, "error parsing source file")
	}

	return nil
}

// Interface implementation
func (bcc *bcc) Parse() error {
	return bcc.parse()
}

// Interface implementation
func (bcc *bcc) Compile() error {
	return bcc.compile()
}

// Interface implementation
func (bcc *bcc) String() string {
	var s string
	for _, inst := range bcc.instructions {
		for _, tok := range inst.tokens {
			s = s + tok.tkn + " "
		}
		s = strings.TrimSuffix(s, " ") + "\n"
	}
	return s
}
