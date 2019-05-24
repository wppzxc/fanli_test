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
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/wpp/fanli_test/pkg/app"
	"github.com/wpp/fanli_test/pkg/process"
	"github.com/wpp/fanli_test/pkg/types"
	"os"
)

var uname string

//var password string
//var toEmail string
var toWeChat string
var duration int64

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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
		conf := types.Config{
			Uname: uname,
			//Password: password,
			//FromEmail: fromEmail,
			//ToEmail: toEmail,
			ToWeChat: toWeChat,
			Duration: duration,
		}
		if err := process.ValidateFlags(conf); err != nil {
			glog.Error(err)
			//fmt.Println(err)
			return
		}
		glog.Infof("The config is : %#v", conf)
		app.AppRun(conf)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glog.Error(err)
		//fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.Flags().StringVarP(&uname, "uname", "u", "uname@163.com", "发送邮件的用户名")
	//rootCmd.Flags().StringVarP(&password, "password", "p", "password", "发送邮件的用户密码")
	//rootCmd.Flags().StringVarP(&fromEmail, "fromEmail", "f", "", "发送邮件的邮箱地址")
	//rootCmd.Flags().StringVarP(&toEmail, "toEmail", "t", "toemail@qq.com", "目的邮箱地址")
	rootCmd.Flags().StringVarP(&uname, "uname", "u", "", "微信openid")
	rootCmd.Flags().StringVarP(&toWeChat, "toWeChat", "w", "user", "目的微信用户")
	rootCmd.Flags().Int64VarP(&duration, "duration", "d", 5, "刷新库存的间隔")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".fanli_test" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".fanli_test")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
