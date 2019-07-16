package utils

import (
	"fmt"
	"github.com/lxn/win"
	"github.com/spf13/pflag"
	"github.com/wpp/fanli_test/pkg/types"
	"k8s.io/klog"
	"strings"
	"time"
)

const (
	msgFormat = `拼多多免单来袭！马上登陆查看！

%s

开始时间: %s

免单网址：http://t.cn/E6KHo26
抢免单，还赚钱，赶快联系客服加入吧！`
	timeFormat = "2006-01-02 15:04:05"
)

func GetMsg(item types.Item) string {
	str := item.ExtendDocument
	str = strings.Replace(str, "#", "", -1)
	str = fmt.Sprintf(msgFormat, str, time.Unix(item.StartTime, 0).Format(timeFormat))
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
	klog.V(9).Infof("Get oldItems : %v", oldItems)
	klog.V(9).Infof("Get newItems : %v", newItems)
	//result := newItems
	for _, o := range oldItems {
		for i, n := range newItems {
			if o.Id == n.Id {
				newItems = append(newItems[:i], newItems[i+1:]...)
				break
			}
		}
	}
	klog.V(9).Infof("Return resultItems : %v", newItems)
	return newItems
}

func ValidateFlags(conf types.Config) error {
	if len(conf.ToUsers) == 0 {
		return fmt.Errorf("Error, toUser is invalidate ")
	}
	if len(conf.Uname) == 0 {
		return fmt.Errorf("Error, uname is invalidate ")
	}
	return nil
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
