package main

import (
	"fmt"
	"github.com/wpp/fanli_test/others/dataoke/cmd"
	"github.com/wpp/fanli_test/pkg/utils"
	"os"
)

func main() {
	command := cmd.NewRootCommand()
	command.Flags().SetNormalizeFunc(utils.WordSepNormalizeFunc)
	command.Flags().MarkHidden("logtostderr")
	command.Flags().MarkHidden("stderrthreshold")
	command.Flags().MarkHidden("vmodule")
	command.Flags().MarkHidden("log-backtrace-at")
	command.Flags().MarkHidden("log-dir")
	command.Flags().MarkHidden("logtostderr")
	command.Flags().MarkHidden("skip-headers")
	command.Flags().MarkHidden("skip-log-headers")
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
