// package main defines the executable for the bcc (bit code compiler) compiler.
package main

import (
	"fmt"
	"os"

	"github.com/mkenney/8bit-cpu/cmd2/pkg/bcc"

	"github.com/bdlm/log/v2"
)

func init() {
	//log.SetFormatter(&log.TextFormatter{DisableTTY: true})
	log.SetLevel(log.DebugLevel)
}

func main() {
	var err error

	sourceFile := os.Args[1]
	destFile := os.Args[2]

	prg, err := bcc.New(sourceFile, destFile)
	if nil != err {
		logger.WithError(err).Fatal("failed to initialize bit code compiler")
	}

	prg.Parse()
	if nil != err {
		logger.WithError(err).Fatal("failed to parse source file")
	}

	err = prg.Compile()
	if nil != err {
		logger.WithError(err).Fatal("failed to compile rom images")
	}

	logger.Info("success")

	code := ""
	for a := range prg.Lines {
		if l, ok := prg.Code[a]; ok {
			code = code + "\n" + l
		}
	}
	fmt.Printf("\n%s\n\n", code)
}
