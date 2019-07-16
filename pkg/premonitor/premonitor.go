package premonitor

import (
	"github.com/wpp/fanli_test/pkg/history"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"time"
)

const (
	preItemsFile = "./preItemsFile"
)

func StartPremonitor(conf types.Config, token string) {
	historyItems := history.GetHistoryItems(preItemsFile)
	defer func(){
		history.UpdateHistoryItems(historyItems)
		if err := history.WriteHistoryItems(historyItems, preItemsFile); err != nil {
			klog.Errorf("Error in write %s", preItemsFile)
		}
	}()
	result, err := utils.GetItems(token, utils.GetPremonitorUrl)
	if err != nil {
		klog.Errorf("Error in get premonitor items : %s", err)
		return
	}
	if result.Count > 0 {
		klog.Info("There found some premonitor items !")
		
		diffItems := utils.GetDiffItems(historyItems, result.Data)
		if len(diffItems) == 0 {
			klog.Error("The premonitor items are modify, But there is no new items !")
			return
		}

		// send message to users
		for _, i := range diffItems {
			msg := utils.GetMsg(i)
			if e := utils.SendMessage(msg, conf.ToUsers); e != nil {
				klog.Errorf("Error on send msg : %s", e)
			} else {
				klog.Infof("Success on send msg to users %v !", conf.ToUsers)
				historyItems = append(historyItems, i)
			}
			time.Sleep(200 * time.Millisecond)
		}
	} else {
		klog.Info("There is no items !")
		return
	}
}
