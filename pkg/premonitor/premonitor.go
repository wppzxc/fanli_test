package premonitor

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"io/ioutil"
	"net/http"
)

const (
	GetPremonitorUrl = "http://v2.yituike.com/fans/fans/proxy_goods?state=2&page=1&limit=10"
)

func StartPremonitor(conf types.Config, md5Str string, token string) string {
	result, err := getItems(token)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err2)
		return md5Str
	}
	if result.Count > 0 {
		glog.Info("There found some premonitor items !")
		//fmt.Println("There found some items !")

		dataStr, err2 := json.Marshal(result)
		if err2 != nil {
			glog.Error(err2)
		}
		h := md5.New()
		h.Write(dataStr)
		newMd5Str := hex.EncodeToString(h.Sum(nil))
		if md5Str == newMd5Str {
			glog.Info("The premonitor items are not modify !")
			//fmt.Println("The items are not modify !")
			return md5Str
		}
		ok := utils.CheckWeChat()
		if ok {
			msg := utils.GetMsg(result)
			if e := utils.SendMessage(msg, conf.ToWeChat); e != nil {
				glog.Errorf("Error on send wechat msg : %s", e)
				//fmt.Println("Error on send wechat msg : ", err)
			} else {
				glog.Info("Success on send msg to wechat !")
				//fmt.Println("Success on send msg to wechat !")
			}
		} else {
			glog.Error("Check WeChat health error ! ")
			//fmt.Println("Check WeChat health error ! ")
		}
		return newMd5Str
	} else {
		glog.Info("There is no items !")
		//fmt.Println("There is no items !")
		return md5Str
	}
}

func getItems(token string) (types.ItemResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetPremonitorUrl, nil)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err)
	}
	req.Header.Add("token", token)
	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	//data = []byte(Data)
	result := types.ItemResult{}
	if err = json.Unmarshal(data, &result); err != nil {
		glog.Error(err, "Maybe the token is invalide ! ")
		//fmt.Println(err)
		return types.ItemResult{}, err
	}
	if result.Count == 0 {
		return types.ItemResult{}, fmt.Errorf("Error in get premonitor items : no items found ")
	}
	return result, nil
}
