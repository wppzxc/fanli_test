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
		klog.Errorf("Error in get premonitor items : %s", err)
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
			klog.Error("The premonitor items are not modify !")
			//fmt.Println("The items are not modify !")
			return md5Str, preItems
		}
		klog.Info("The premonitor items changed")
		diffItems := utils.GetDiffItems(preItems, result.Data)
		if len(diffItems) == 0 {
			klog.Error("The premonitor items are modify, But there is no new items !")
			return newMd5Str, result.Data
		}

		// send message to users
		for _, i := range diffItems {
			msg := utils.GetMsg(i)
			if e := utils.SendMessage(msg, conf.ToUsers); e != nil {
				klog.Errorf("Error on send %s msg : %s", conf.Process, e)
			} else {
				klog.Infof("Success on send msg to %s !", conf.Process)
			}
			time.Sleep(200 * time.Millisecond)
		}

		return newMd5Str, result.Data
	} else {
		klog.Info("There is no items !")
		return md5Str, preItems
	}
}
