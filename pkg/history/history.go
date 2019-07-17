package history

import (
	"encoding/json"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"k8s.io/klog"
	"os"
	"path"
	"time"
)

const (
	defaultFile = "./history"
)

func GetHistoryItems(file string) []types.Item {
	if len(file) == 0 {
		file = defaultFile
	}
	_, err := os.Stat(file)
	if os.IsPermission(err) {
		klog.Error(err)
		return nil
	}
	if os.IsNotExist(err) {
		klog.Infof("create the file %s ", file)
		if err := os.MkdirAll(path.Dir(file), 0755); err != nil {
			klog.Error(err)
		}
		_, err := os.Create(file)
		if err != nil {
			klog.Error(err)
		}
		return nil
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		klog.Error(err)
		return nil
	}
	if len(data) == 0 {
		return nil
	}
	var items []types.Item
	if err := json.Unmarshal(data, &items); err != nil {
		klog.Error(err)
		return nil
	}
	return items
}

func UpdateHistoryItems(items []types.Item) {
	now := time.Now().Unix()
	for i := 0; i < len(items); i++ {
		if items[i].StopTime < now {
			items = append(items[:i], items[i+1:]...)
			i--
		}
	}
}

func WriteHistoryItems(items []types.Item, file string) error {
	data, err := json.Marshal(items)
	if err != nil {
		klog.Error(err)
		return err
	}
	if err := ioutil.WriteFile(file, data, 0755); err != nil {
		klog.Error(err)
		return err
	}
	return nil
}
