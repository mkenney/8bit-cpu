package bcc

import (
	"strings"

	"github.com/bdlm/errors"
)

func newToken(ln int, pos int, prt string) (*tok, error) {
	var err error

	tkn := &tok{
		ln:  ln,
		pos: pos,
		tkn: prt,
		typ: TOK_NIL,
	}

	err = tkn.tokenize()
	if nil != err {
		return nil, errors.Wrap(err, "error generating token")
	}

	return tkn, nil
}

func (tkn *tok) tokenize() error {
	var err error

	switch tkn.pos {
	case 0:
		if strings.HasPrefix(tkn.tkn, "$") {
			tkn.typ = TOK_CONST
		} else if strings.HasSuffix(tkn.tkn, "{") {
			tkn.typ = TOK_SUB
		} else if "}" == tkn.tkn {
			tkn.typ = TOK_SUBEND
		} else if "" != tkn.tkn {
			tkn.typ = TOK_LABEL
		}

	case 1:
		// check for operation codes
		if _, ok := opMap[tkn.tkn]; ok {
			tkn.typ = TOK_OP

		} else {
			// check for literals
			if tkn.dat, err = parseLiteral(tkn.tkn); nil != err {
				return errors.Errorf("unknown data literal '%s' at line %d", tkn.tkn, tkn.ln)
			}
			tkn.typ = TOK_LIT
		}

	case 2:
		// check for constant references
		if strings.HasPrefix(tkn.tkn, "$") {
			tkn.typ = TOK_REF

			// check for literals
		} else if tkn.dat, err = parseLiteral(tkn.tkn); nil != err {
			return errors.Wrap(err, "unknown literal '%s' at line %d", tkn.tkn, tkn.ln)
		}
	}

	return nil
}

type tokenType string

const (
	// unknown
	TOK_NIL tokenType = ""
	// defined constant
	TOK_CONST tokenType = "TOK_CONST"
	// constant reference
	TOK_REF tokenType = "TOK_REF"
	// data literal
	TOK_LIT tokenType = "TOK_LIT"
	// jump label
	TOK_LABEL tokenType = "TOK_LABEL"
	// operation
	TOK_OP tokenType = "TOK_OP"
	// subroutine
	TOK_SUB tokenType = "TOK_SUB"
	// subroutine end
	TOK_SUBEND tokenType = "TOK_SUBEND"
)

type tok struct {
	// source file line number
	ln int
	// line position. positions are space delimited. 0, 1, or 2.
	pos int
	// token string
	tkn string
	// token type
	typ tokenType
	// data for literals
	dat byte
}

func (tok *tok) Type() tokenType {
	return tok.typ
}
