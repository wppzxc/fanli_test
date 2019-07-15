// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/wpp/fanli_test/cmd"
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
	command.Flags().MarkHidden("alsologtostderr")
	command.Flags().MarkHidden("skip-headers")
	command.Flags().MarkHidden("skip-log-headers")
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
