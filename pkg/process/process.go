package process

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"time"
)

func StartProcess(conf types.Config, md5Str string, proItems []types.Item, token string) (string, []types.Item) {
	result, err := utils.GetItems(token, utils.GetProcessUrl)
	if err != nil {
		klog.Errorf("Error in get process items : %s", err)
		//fmt.Println(err2)
		return md5Str, proItems
	}
	if result.Count > 0 {
		klog.Info("There found some process items !")
		//fmt.Println("There found some items !")

		dataStr, err2 := json.Marshal(result)
		if err2 != nil {
			klog.Error(err2)
		}
		h := md5.New()
		h.Write(dataStr)
		newMd5Str := hex.EncodeToString(h.Sum(nil))
		if md5Str == newMd5Str {
			klog.Error("The process items are not modify !")
			//fmt.Println("The items are not modify !")
			return md5Str, proItems
		}
		klog.Info("The process items changed")
		diffItems := utils.GetDiffItems(proItems, result.Data)
		if len(diffItems) == 0 {
			klog.Error("The process items are modify but there is no new items !")
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
		return md5Str, proItems
	}
}
