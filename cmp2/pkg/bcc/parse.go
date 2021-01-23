package bcc

import (
	"strconv"
	"strings"

	"github.com/bdlm/errors"
)

// parseLiteral
func parseLiteral(dataStr string) (byte, error) {
	var err error
	var data int64

	switch true {
	// binary
	case strings.HasPrefix(dataStr, "0b"):
		data, err = strconv.ParseInt(string(dataStr[2:]), 2, 8)
	// hex
	case strings.HasPrefix(dataStr, "0x"):
		data, err = strconv.ParseInt(string(dataStr[2:]), 16, 8)
	// decimal
	default:
		data, err = strconv.ParseInt(dataStr, 10, 8)
	}
	if nil != err {
		return 0, errors.Wrap(err, "failure parsing data '%s'", dataStr)
	}

	return byte(data), nil
}
