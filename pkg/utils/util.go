package utils

import (
	"fmt"
	"github.com/lxn/win"
	"github.com/spf13/pflag"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"k8s.io/klog"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
	"time"
)

const (
	msgFormat = `预告中有新增拼多多免单！欢迎查看！
拼多多免单来袭！马上登陆查看！

%s

开始时间: %s

免单网址：%s`
	timeFormat = "2006-01-02 15:04:05"
)

func GetMsg(item types.Item, link string) string {
	str := item.ExtendDocument
	str = strings.Replace(str, "#", "", -1)
	str = fmt.Sprintf(msgFormat, str, time.Unix(item.StartTime, 0).Format(timeFormat), link)
	return str
}

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}

func GetDiffItems(oldItems []types.Item, newItems []types.Item) []types.Item {
	klog.V(9).Infof("Get historyItems : %v", oldItems)
	klog.V(9).Infof("Get newItems : %v", newItems)
	//result := newItems
	for _, o := range oldItems {
		for i := 0; i < len(newItems); i++ {
			if o.ExtendDocument == newItems[i].ExtendDocument && o.StartTime == newItems[i].StartTime {
				newItems = append(newItems[:i], newItems[i+1:]...)
				i--
				break
			}
		}
	}
	klog.V(9).Infof("Return diffItems : %v", newItems)
	return newItems
}

func ValidateConfig(file string) (*types.Config, error) {
	_, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Error in get data from config file ")
	}
	config := new(types.Config)
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("Error in decode config file %s ", err)
	}
	if len(config.Auth.Username) == 0 || len(config.Auth.Password) == 0 {
		return nil, fmt.Errorf("Error in config file, username or password must provide ")
	}
	if !config.Fanli.Process.Start && !config.Fanli.Premonitor.Start {
		return nil, fmt.Errorf("Error in config file, process or premonitor must start one ")
	}
	if len(config.Receivers) == 0 {
		return nil, fmt.Errorf("Error in config file, receiver must provide ")
	}
	if config.Fanli.RefreshInterval == 0 {
		klog.Info("The fanli interval get 0, set to default 120 ")
		config.Fanli.RefreshInterval = 120
	}
	if config.Fanli.SendInterval == 0 {
		klog.Info("The fanli interval get 0, set to default 120 ")
		config.Fanli.SendInterval = 1
	}
	klog.Info("validate config file ok")
	return config, nil
}

func SetForegroundWindow(hWnd win.HWND) bool {
	hForeWnd := win.GetForegroundWindow()
	dwCurID := win.GetCurrentThreadId()
	dcID := int32(dwCurID)
	dwForeID := win.GetWindowThreadProcessId(hForeWnd, nil)
	dfID := int32(dwForeID)
	win.AttachThreadInput(dcID, dfID, true)
	win.ShowWindow(hWnd, win.SW_SHOWNORMAL)
	win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE)
	win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOSIZE|win.SWP_NOMOVE)
	ok := win.SetForegroundWindow(hWnd)
	win.AttachThreadInput(dcID, dfID, false)
	return ok
}
