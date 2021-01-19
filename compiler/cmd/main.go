// package main defines the executable for the bcc (bit code compiler) compiler.
package main

import (
	"encoding/json"
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

	log.Debug("initializing program")
	prg, err := program.New(os.Args[1])
	if nil != err {
		log.WithError(err).Fatal("failed to initialize program parser")
	}

	log.Debug("compiling...")
	err = prg.Compile()
	if nil != err {
		log.WithError(err).Fatal("failed to compile binary")
	}

	b, _ := json.MarshalIndent(prg.Code, "", "\t")
	fmt.Println(string(b))

	log.WithField("tokens", string(b)).Info("success")
}
