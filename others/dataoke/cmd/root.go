package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wpp/fanli_test/others/dataoke/app"
	appVersion "github.com/wpp/fanli_test/pkg/version"
	"k8s.io/klog"
	"os"
)

var (
	version bool
	quan    bool
	top     bool
	begin   int
	end     int
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
				fmt.Print("fanli.exe : ", appVersion.Get())
				os.Exit(0)
			}
			if quan {
				klog.Info("抓取领券直播数据...")
				if err := app.StartQuan(begin, end); err != nil {
					klog.Error(err)
					os.Exit(1)
				}
			}
			if top {
				klog.Info("抓取实时榜单数据...")
				if err := app.StartTop(); err != nil {
					klog.Error(err)
					os.Exit(1)
				}
			}
			klog.Info("抓取成功")
			os.Exit(0)
		},
	}
	rootCmd.Flags().BoolVar(&version, "version", false, "The version of fanli.exe")
	rootCmd.Flags().BoolVar(&quan, "quan", false, "领券直播")
	rootCmd.Flags().BoolVar(&top, "top", false, "实时榜单")
	rootCmd.Flags().IntVarP(&begin, "begin", "b", 1, "开始页")
	rootCmd.Flags().IntVarP(&end, "end", "e", 1, "结束页")
	klogFlagset := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlagset)
	rootCmd.Flags().AddGoFlagSet(klogFlagset)
	return rootCmd
}
