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
	"github.com/spf13/cobra"
	"github.com/wpp/fanli_test/pkg/app"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
)

func NewRootCommand() *cobra.Command {

	conf := types.Config{}
	rootCmd := &cobra.Command{
		Use:   "fanli_test",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.ValidateFlags(conf); err != nil {
				klog.Error(err)
				//fmt.Println(err)
				return
			}
			klog.Infof("The config is : %#v", conf)
			app.AppRun(conf)
		},
	}

	rootCmd.Flags().StringVarP(&conf.Uname, "uname", "u", "", "微信openid")
	rootCmd.Flags().StringVarP(&conf.Process, "process", "p", "WeChat", "目标程序，WeChat 或 TIM 或 腾讯QQ")
	rootCmd.Flags().StringVarP(&conf.ToWeChat, "toUser", "t", "", "目的用户")
	rootCmd.Flags().Int64VarP(&conf.Duration, "duration", "d", 5, "刷新库存的间隔")
	klogFlagset := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlagset)
	rootCmd.Flags().AddGoFlagSet(klogFlagset)
	return rootCmd
}
