package app

import (
	"encoding/json"
	"fmt"
	"github.com/wpp/fanli_test/pkg/premonitor"
	"github.com/wpp/fanli_test/pkg/process"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"time"
)

const (
	GetTokenUrlPrefix = "http://v2.yituike.com/admin/Weixinm/login?access_token=098f6bcd4621d373cade4e832627b4f6&openid="
	//os_Ph0j7eHfGJhowF8E-_kJc1fiM
//	-u os_Ph0j7eHfGJhowF8E-_kJc1fiM -w bigben -d 5
)

func AppRun(conf types.Config) {
	klog.Info("Start process ... !")
	//fmt.Println("Start process ... !")
	token, err := getToken(conf)
	if err != nil {
		klog.Fatal(err)
	}
	md5ProStr := "test"
	proItems := []types.Item{}
	md5PreStr := "test"
	preItems := []types.Item{}
	for range time.Tick(time.Duration(conf.Duration) * time.Second) {
		md5ProStr, proItems, md5PreStr, preItems = start(conf, md5ProStr, proItems, md5PreStr, preItems, token)
	}
}

func start(conf types.Config, md5ProStr string, proItems []types.Item, md5PreStr string, preItems []types.Item, token string) (string, []types.Item, string, []types.Item) {
	klog.V(9).Infof("input proItems : %v", proItems)
	klog.V(9).Infof("input preItems : %v", preItems)
	md5ProStrNew, newProItems := process.StartProcess(conf, md5ProStr, proItems, token)
	md5PreStrNew, newPreItems := premonitor.StartPremonitor(conf, md5PreStr, preItems, token)
	klog.V(9).Infof("return proItems : %v", newProItems)
	klog.V(9).Infof("return preItems : %v", newPreItems)
	return md5ProStrNew, newProItems, md5PreStrNew, newPreItems
}

func getToken(conf types.Config) (string, error) {
	if len(conf.Uname) == 0 {
		return "", fmt.Errorf("invalide uname ! : %s ", conf.Uname)
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(GetTokenUrlPrefix, conf.Uname, "&invite_id=007"), nil)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := types.TokenResult{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	if result.Error != 0 {
		return "", fmt.Errorf("Error in get token : %s ", err)
	}
	return result.Token, nil
}
