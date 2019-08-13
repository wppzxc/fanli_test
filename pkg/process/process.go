package process

import (
	"github.com/wpp/fanli_test/pkg/history"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"k8s.io/klog"
	"os"
	"time"
)

const (
	proItemsFile = "./proItemsFile"
)

type Processer struct {
	Config *types.Config
}

func (p *Processer) StartProcess() {
	historyItems := history.GetHistoryItems(proItemsFile)
	klog.Info("Start process ...")

	defer func() {
		history.UpdateHistoryItems(historyItems)
		if err := history.WriteHistoryItems(historyItems, proItemsFile); err != nil {
			klog.Errorf("Error in write %s", proItemsFile)
		}
	}()

	token, err := utils.GetToken(&p.Config.Auth)
	if err != nil {
		klog.Error(err)
		return
	}
	result, err := utils.GetItems(token, p.Config.Fanli.Process.Url)
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

		// only send first item msg, so mark all diffItems to historyItems
		for _, i := range diffItems {
			historyItems = append(historyItems, i)
		}

		tmpfile, err := utils.SaveImage(diffItems[0].GoodsImageUrl)
		if err != nil {
			klog.Errorf("Error in download image %s", err)
		}
		defer os.Remove(tmpfile.Name())

		// send message to users
		for _, u := range p.Config.Receivers {
			if err := utils.SendImage(tmpfile, u); err != nil {
				klog.Errorf("Error in send image to user %s", err)
			} else {
				klog.Infof("Success on send image to user %s ", u.Name)
			}

			msg := utils.GetMsg(diffItems[0], u.Link)
			if err := utils.SendMessage(msg, u.Name); err != nil {
				klog.Errorf("Error on send msg : %s", err)
			} else {
				klog.Infof("Success on send msg to users %s ", u.Name)
			}
			time.Sleep(time.Duration(p.Config.Fanli.SendInterval) * time.Second)
		}
	} else {
		klog.Info("There is no items !")
		return
	}
}
