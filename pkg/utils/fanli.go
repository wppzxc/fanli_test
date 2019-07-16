package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
)

const (
	GetProcessUrl    = "http://v2.yituike.com/fans/fans/proxy_goods?state=1&page=1&limit=10"
	GetPremonitorUrl = "http://v2.yituike.com/fans/fans/proxy_goods?state=2&page=1&limit=10"
)

func GetItems(token string, url string) (types.ItemResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		klog.Error(err)
	}
	req.Header.Add("token", token)
	resp, err := client.Do(req)
	if err != nil {
		klog.Error(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := types.ItemResult{}
	if err = json.Unmarshal(data, &result); err != nil {
		klog.Errorf("Maybe the token is invalide !  result : %s, error : %s", data, err)
		return types.ItemResult{}, err
	}
	if result.Count == 0 {
		return types.ItemResult{}, fmt.Errorf("Error in get premonitor items : no items found ")
	}
	return result, nil
}
