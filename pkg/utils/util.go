package utils

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/wpp/fanli_test/pkg/types"
	"k8s.io/klog"
	"strings"
	"time"
)

const (
	msgFormat = `拼多多免单来袭！马上登陆查看！

%s

开始时间: %s`
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
	result := newItems
	for _, o := range oldItems {
		for i, n := range newItems {
			if o.Id == n.Id {
				result = append(result[:i], result[i+1:]...)
				break
			}
		}
	}
	klog.V(9).Infof("Return resultItems : %v", result)
	return result
}

func ValidateFlags(conf types.Config) error {
	if len(conf.ToWeChat) == 0 {
		return fmt.Errorf("Error, toUser is invalidate ")
	}
	if len(conf.Uname) == 0 {
		return fmt.Errorf("Error, uname is invalidate ")
	}
	return nil
}
