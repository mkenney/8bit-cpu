// package main defines the executable for the bcc (bit code compiler) compiler.
package main

import (
	"os"

	"github.com/mkenney/8bit-cpu/cmp2/pkg/bcc"

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

	logger := log.WithFields(log.Fields{"src": os.Args[1], "dest": os.Args[2]})
	logger.Debug("initializing compiler")
	prg, err := bcc.New(sourceFile, destFile)
	if nil != err {
		logger.WithError(err).Fatal("failed to initialize bit code compiler")
	}

	logger.Debug("parsing src file")
	err = prg.Parse()
	if nil != err {
		logger.WithError(err).Fatal("failed to parse source file")
	}

	logger.Debug("writing dest image")
	err = prg.Compile()
	if nil != err {
		logger.WithError(err).Fatal("failed to compile ROM images")
	}

	logger.Info("success")

	// DEBUG
	//code := ""
	//for a := range prg.Lines() {
	//	l := prg.Instructions()[a]
	//	if l, ok := prg.Instructions()[a]; ok {
	//		code = code + "\n" + l
	//	}
	//}
	//fmt.Printf("\n%s\n\n", code)
}
