package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"net/url"
)

func GetItems(token string, url string) (types.ItemResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		klog.Error(err)
	}
	//req.Header.Add("token", token)
	req.Header.Set("token", token)
	resp, err := client.Do(req)
	if err != nil {
		klog.Error(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := types.ItemResult{}
	if err = json.Unmarshal(data, &result); err != nil {
		if result.Count == "" {
			err = fmt.Errorf("There is no goods found ")
		}
		return types.ItemResult{}, err
	}
	if result.Count == "0" {
		return types.ItemResult{}, fmt.Errorf("no items found ")
	}
	return result, nil
}

func GetToken(auth *types.AuthInfo) (string, error) {
	
	data := make(url.Values)
	data["account"] = []string{auth.Username}
	data["password"] = []string{auth.Password}
	resp, err := http.PostForm(auth.Url, data)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	defer resp.Body.Close()
	d, _ := ioutil.ReadAll(resp.Body)
	result := new(types.TokenResult)
	err = json.Unmarshal(d, result)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	if result.Code != 0 {
		return "", fmt.Errorf("Error in get token : %s ", err)
	}
	return result.Data[0].Token, nil
}
