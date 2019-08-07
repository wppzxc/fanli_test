package premonitor

import (
	"github.com/wpp/fanli_test/pkg/history"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"strconv"
	"time"
)

const (
	preItemsFile = "./preItemsFile"
)

type Premonitor struct {
	Config *types.Config
}

func (p *Premonitor) StartPremonitor() {
	klog.Info("start premonitor... ")
	historyItems := history.GetHistoryItems(preItemsFile)
	
	defer func(){
		history.UpdateHistoryItems(historyItems)
		if err := history.WriteHistoryItems(historyItems, preItemsFile); err != nil {
			klog.Errorf("Error in write %s", preItemsFile)
		}
	}()
	
	token, err := utils.GetToken(&p.Config.Auth)
	result, err := utils.GetItems(token, p.Config.Fanli.Premonitor.Url)
	if err != nil {
		klog.Errorf("Error in get premonitor items : %s", err)
		return
	}
	count, err := strconv.Atoi(result.Count)
	if err != nil {
		klog.Errorf("Error in decode result.Count to int %s", result.Count)
		return
	}
	if count > 0 {
		klog.Info("There found some premonitor items !")
		
		diffItems := utils.GetDiffItems(historyItems, result.Data)
		if len(diffItems) == 0 {
			klog.Error("The premonitor items are modify, But there is no new items !")
			return
		}
		
		// only send first item msg, so mark all diffItems to historyItems
		for _, i := range diffItems {
			historyItems = append(historyItems, i)
		}

		// send message to users
		for _, u := range p.Config.Receiver {
			msg := utils.GetMsg(diffItems[0], u.Link)
			if err := utils.SendMessage(msg, u.Name); err != nil {
				klog.Errorf("Error on send msg : %s", err)
			} else {
				klog.Infof("Success on send msg to users %v !", u.Name)
			}
			time.Sleep(200 * time.Millisecond)
		}
	} else {
		klog.Info("There is no items !")
		return
	}
}
