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

package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wpp/fanli_test/pkg/app"
	"github.com/wpp/fanli_test/pkg/premonitor"
	"github.com/wpp/fanli_test/pkg/process"
	"github.com/wpp/fanli_test/pkg/utils"
	appVersion "github.com/wpp/fanli_test/pkg/version"
	"k8s.io/klog"
	"os"
)

var (
	version bool
	config  string
)

func NewRootCommand() *cobra.Command {
	
	rootCmd := &cobra.Command{
		Use:   "fanli.exe",
		Short: "Used to get process or premonitor items and send it to users",
		Long:  `There is no description of the fanli.exe`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				fmt.Print("fanli.exe : %s ", appVersion.Get())
				os.Exit(0)
			}
			conf, err := utils.ValidateConfig(config)
			if err != nil {
				klog.Fatal(err)
			}
			fs := []func(){}
			if conf.Fanli.Process.Start {
				pro := process.Processer{
					Config: conf,
				}
				klog.Info("register process ")
				fs = append(fs, pro.StartProcess)
			}
			if conf.Fanli.Premonitor.Start {
				pre := premonitor.Premonitor{
					Config: conf,
				}
				klog.Info("register premonitor ")
				fs = append(fs, pre.StartPremonitor)
			}
			app.AppRun(conf, fs)
		},
	}
	rootCmd.Flags().BoolVar(&version, "version", false, "The version of fanli.exe")
	rootCmd.Flags().StringVar(&config, "config", "", "The config file of fanli.exe")
	klogFlagset := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlagset)
	rootCmd.Flags().AddGoFlagSet(klogFlagset)
	return rootCmd
}
