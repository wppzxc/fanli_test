package process

import (
	"github.com/wpp/fanli_test/pkg/history"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"time"
)

const (
	proItemsFile = "./proItemsFile"
)

func StartProcess(conf types.Config, token string) {
	historyItems := history.GetHistoryItems(proItemsFile)
	defer func(){
		history.UpdateHistoryItems(historyItems)
		if err := history.WriteHistoryItems(historyItems, proItemsFile); err != nil {
			klog.Errorf("Error in write %s", proItemsFile)
		}
	}()
	result, err := utils.GetItems(token, utils.GetProcessUrl)
	if err != nil {
		klog.Errorf("Error in get process items : %s", err)
		return
	}
	if result.Count > 0 {
		klog.Info("There found some process items !")

		diffItems := utils.GetDiffItems(historyItems, result.Data)
		if len(diffItems) == 0 {
			klog.Error("The process items are modify but there is no new items !")
			return
		}

		// send message to users
		for _, i := range diffItems {
			msg := utils.GetMsg(i)
			if e := utils.SendMessage(msg, conf.ToUsers); e != nil {
				klog.Errorf("Error on send %s msg : %s", conf.Process, e)
			} else {
				klog.Infof("Success on send msg to %s !", conf.Process)
				historyItems = append(historyItems, i)
			}
			time.Sleep(200 * time.Millisecond)
		}
	} else {
		klog.Info("There is no items !")
		return
	}
}
