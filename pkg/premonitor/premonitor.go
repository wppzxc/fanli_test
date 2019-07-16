package premonitor

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"time"
)

func StartPremonitor(conf types.Config, md5Str string, preItems []types.Item, token string) (string, []types.Item) {
	result, err := utils.GetItems(token, utils.GetPremonitorUrl)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err2)
		return md5Str, preItems
	}
	if result.Count > 0 {
		klog.Info("There found some premonitor items !")
		//fmt.Println("There found some items !")

		dataStr, err2 := json.Marshal(result)
		if err2 != nil {
			klog.Error(err2)
		}
		h := md5.New()
		h.Write(dataStr)
		newMd5Str := hex.EncodeToString(h.Sum(nil))
		if md5Str == newMd5Str {
			klog.Warning("The premonitor items are not modify !")
			//fmt.Println("The items are not modify !")
			return md5Str, preItems
		}
		klog.Info("The premonitor items changed")
		ok := utils.CheckProcess(conf.Process)
		diffItems := utils.GetDiffItems(preItems, result.Data)
		if len(diffItems) == 0 {
			klog.Warning("But there is no new items !")
			return newMd5Str, result.Data
		}
		if ok {
			for _, i := range diffItems {
				msg := utils.GetMsg(i)
				if e := utils.SendMessage(msg, conf.ToWeChat); e != nil {
					klog.Errorf("Error on send %s msg : %s", conf.Process, e)
				} else {
					klog.Infof("Success on send msg to %s !", conf.Process)
				}
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			klog.Errorf("Check %s health error ! ", conf.Process)
			//fmt.Println("Check WeChat health error ! ")
		}
		return newMd5Str, result.Data
	} else {
		klog.Info("There is no items !")
		//fmt.Println("There is no items !")
		return md5Str, preItems
	}
}
