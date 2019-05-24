package app

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/wpp/fanli_test/pkg/premonitor"
	"github.com/wpp/fanli_test/pkg/process"
	"github.com/wpp/fanli_test/pkg/types"
	"io/ioutil"
	"net/http"
	"time"
)
const (
	GetTokenUrl = "http://v2.yituike.com/admin/Weixinm/login?access_token=test&openid=os_Ph0hGJxQNrHmNWrL5xV3nuTuc&invite_id=208"
)

func AppRun(conf types.Config) {
	glog.Info("Start process ... !")
	//fmt.Println("Start process ... !")
	token, err := getToken()
	if err != nil {
		glog.Error(err)
	}
	md5ProStr := "test"
	md5PreStr := "test"
	for range time.Tick(5 * time.Second) {
		md5ProStr, md5PreStr = start(conf, md5ProStr, md5PreStr, token)
	}
}

func start(conf types.Config, md5ProStr string, md5PreStr string, token string) (string, string) {
	md5ProStrNew := process.StartProcess(conf, md5ProStr, token)
	md5PreStrNew := premonitor.StartPremonitor(conf, md5PreStr, token)
	return md5ProStrNew, md5PreStrNew
}

func getToken() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetTokenUrl, nil)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	data,_ := ioutil.ReadAll(resp.Body)
	result := types.TokenResult{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err)
	}
	if result.Error != 0 {
		return "", fmt.Errorf("Error in get token : %s ", err)
	}
	return result.Token, nil
}
