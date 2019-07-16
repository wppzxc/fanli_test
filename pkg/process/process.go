package process

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wpp/fanli_test/pkg/types"
	"github.com/wpp/fanli_test/pkg/utils"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"time"
)

const (
	GetProcessUrl = "http://v2.yituike.com/fans/fans/proxy_goods?state=1&page=1&limit=10"
	Data          = `{
    "code": 0,
    "msg": "",
    "count": 1,
    "data": [
        {
            "id": 2581,
            "extend_document": "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
            "goods_image_url": "http://m.wangshikun.wang/2019_7_15_211349微信图片_20190715211150.jpg?imageMogr2/auto-orient/thumbnail/300x300/format/jpg/blur/1x0/quality/75|imageslim",
            "goods_name": "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
            "refund_amount": "2.80",
            "min_group_price": "3.80",
            "coupon_discount": "1.00",
            "start_time": 1563197400,
            "stop_time": 1563283800
        }
    ]
}`
)

func ValidateFlags(conf types.Config) error {
	if len(conf.ToWeChat) == 0 {
		return fmt.Errorf("Error, toWeChat is invalidate ")
	}
	if len(conf.Uname) == 0 {
		return fmt.Errorf("Error, uname is invalidate ")
	}
	return nil
}

func StartProcess(conf types.Config, md5Str string, proItems []types.Item, token string) (string, []types.Item) {
	result, err := getItems(token)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err2)
		return md5Str, proItems
	}
	if result.Count > 0 {
		klog.Info("There found some process items !")
		//fmt.Println("There found some items !")

		dataStr, err2 := json.Marshal(result)
		if err2 != nil {
			klog.Error(err2)
		}
		h := md5.New()
		h.Write(dataStr)
		newMd5Str := hex.EncodeToString(h.Sum(nil))
		if md5Str == newMd5Str {
			klog.Info("The process items are not modify !")
			//fmt.Println("The items are not modify !")
			return md5Str, proItems
		}
		klog.Info("The process items changed")
		ok := utils.CheckWeChat()
		diffItems := utils.GetDiffItems(proItems, result.Data)
		if len(diffItems) == 0 {
			klog.Info("But there is no new items !")
			klog.V(9).Info("oldItems : ", proItems)
			klog.V(9).Info("newItems : ", result.Data)
			return newMd5Str, result.Data
		}
		if ok {
			for _, i := range diffItems {
				msg := utils.GetMsg(i)
				if e := utils.SendMessage(msg, conf.ToWeChat); e != nil {
					klog.Errorf("Error on send wechat msg : %s", e)
				} else {
					klog.Info("Success on send msg to wechat !")
				}
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			klog.Error("Check WeChat health error ! ")
			//fmt.Println("Check WeChat health error ! ")
		}
		return newMd5Str, result.Data
	} else {
		klog.Info("There is no items !")
		//fmt.Println("There is no items !")
		return md5Str, proItems
	}
}

func getItems(token string) (types.ItemResult, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetProcessUrl, nil)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	req.Header.Add("token", token)
	resp, err := client.Do(req)
	if err != nil {
		klog.Error(err)
		//fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	result := types.ItemResult{}
	if err = json.Unmarshal(data, &result); err != nil {
		klog.Errorf("Maybe the token is invalide ! result : %s, error : %s", data, err)
		//fmt.Println(err)
		return types.ItemResult{}, err
	}
	if result.Count == 0 {
		return types.ItemResult{}, fmt.Errorf("Error in get process items : no items found ")
	}
	return result, nil
}
