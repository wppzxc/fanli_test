package utils

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/wpp/fanli_test/pkg/types"
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
	result := newItems
	for _, o := range oldItems {
		for i, n := range newItems {
			if o.Id == n.Id {
				result = append(result[:i], result[i+1:]...)
				break
			}
		}
	}
	return result
}
