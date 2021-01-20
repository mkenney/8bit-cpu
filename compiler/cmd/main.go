// package main defines the executable for the bcc (bit code compiler) compiler.
package main

import (
	"fmt"
	"os"

	"github.com/mkenney/8bit-cpu/compiler/pkg/program"

	"github.com/bdlm/log/v2"
)

func init() {
	log.SetFormatter(&log.TextFormatter{DisableTTY: true})
	log.SetLevel(log.DebugLevel)
}

func main() {
	var err error
	logger := log.WithField("src", os.Args[1])

	logger.Info("parsing...")
	prg, err := program.New(os.Args[1])
	if nil != err {
		logger.WithError(err).
			Fatal("failed to parse source file")
	}

	dest := os.Args[1] + ".bin"
	logger = logger.WithField("dest", dest)
	logger.Info("compiling...")
	err = prg.Compile(dest)
	if nil != err {
		logger.WithError(err).
			Fatal("failed to compile binary")
	}
	logger.Info("success")

	code := ""
	for a := range prg.Lines {
		if l, ok := prg.Code[a]; ok {
			code = code + "\n" + l
		}
	}
	fmt.Printf("\n\n%s\n\n", code)
	//b, _ := json.MarshalIndent(prg.Code, "", "\t")
	//fmt.Printf("\n\n%s\n\n", strings.Join(prg.Lines, "\n"))
}
