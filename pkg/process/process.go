package process

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
	GetProcessUrl = "http://v2.yituike.com/fans/fans/proxy_goods?state=1&page=1&limit=10"
	Data = `{
   "code": 0,
   "msg": "",
   "count": 2,
   "data": [
       {
           "id": 474,
           "extend_document": "###1张老鼠贴免单来袭\n###拍  3张装付款4.5元（实际发一张老鼠贴）\n###券无代表结束，价格不对不要拍\n### 北京河北地区的  名字需要真实姓名 拍，快递要实名制，不然不返款（切记）\n###有事找群主",
           "announce_document": "###1张老鼠贴免单来袭\n###拍  3张装付款4.5元（实际发一张老鼠贴）\n###券无代表结束，价格不对不要拍\n### 北京河北地区的  名字需要真实姓名 拍，快递要实名制，不然不返款（切记）\n###有事找群主",
           "goods_image_url": "http://m.wangshikun.wang/2019_5_22_193117粘鼠板一张1.jpg?imageMogr2/auto-orient/thumbnail/300x300/format/jpg/blur/1x0/quality/75|imageslim",
           "goods_name": "老鼠贴粘鼠板超强力灭鼠电猫捕鼠神器驱鼠器老鼠板老鼠夹捕鼠笼",
           "refund_amount": "4.50",
           "min_group_price": "8.50",
           "coupon_discount": "4.00",
           "start_time": 1558454400,
           "stop_time": 1558540800
       },
       {
           "id": 469,
           "extend_document": "抢免单了！抢免单了！\n拍【5片10贴缓解眼部疲劳】拍【6份】【付款价格28.7】元\n发货眼贴3袋\n优惠券【1】元没有券代表活动结束，价格不对不要拍\n确认收货五星好评之后，48小时之内自动返款到用户链接端账户余额，点击提现即可返款\n一个人只能拍一次，严禁一人多个账号，重复下单，否则不返款\n任何问题联系群主解决，不要联系商家\n商家不易感谢配合",
           "announce_document": "抢免单了！抢免单了！\n拍【5片10贴缓解眼部疲劳】拍【6份】【付款价格28.7】元\n发货眼贴3袋\n优惠券【1】元没有券代表活动结束，价格不对不要拍\n确认收货五星好评之后，48小时之内自动返款到用户链接端账户余额，点击提现即可返款\n一个人只能拍一次，严禁一人多个账号，重复下单，否则不返款\n任何问题联系群主解决，不要联系商家\n商家不易感谢配合",
           "goods_image_url": "http://m.wangshikun.wang/2019_5_22_174418.....jpg?imageMogr2/auto-orient/thumbnail/300x300/format/jpg/blur/1x0/quality/75|imageslim",
           "goods_name": "【7天淡化黑眼圈】海藻眼膜贴去皱纹5片眼霜眼贴黑眼圈去眼袋眼纹",
           "refund_amount": "28.70",
           "min_group_price": "29.70",
           "coupon_discount": "1.00",
           "start_time": 1558524600,
           "stop_time": 1558537200
       }
   ]
}`
)

func ValidateFlags(conf types.Config) error {
	if len(conf.ToWeChat) == 0 {
		return fmt.Errorf("Error, %s is invalidate ", conf.ToWeChat)
	}
	return nil
}

func StartProcess(conf types.Config, md5Str string, token string) string {
	result, err := getItems(token)
	if err != nil {
		glog.Error(err)
		//fmt.Println(err2)
		return md5Str
	}
	if result.Count > 0 {
		glog.Info("There found some process items !")
		//fmt.Println("There found some items !")

		dataStr, err2 := json.Marshal(result)
		if err2 != nil {
			glog.Error(err2)
		}
		h := md5.New()
		h.Write(dataStr)
		newMd5Str := hex.EncodeToString(h.Sum(nil))
		if md5Str == newMd5Str {
			glog.Info("The process items are not modify !")
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
	req, err := http.NewRequest(http.MethodGet, GetProcessUrl, nil)
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
	data,_ := ioutil.ReadAll(resp.Body)
	//data = []byte(Data)
	result := types.ItemResult{}
	err = json.Unmarshal(data, &result)
	//items := []types.Item{}
	if err != nil {
		glog.Error(err)
		//fmt.Println(err)
		return types.ItemResult{}, err
	}
	if result.Count == 0 {
		return types.ItemResult{}, fmt.Errorf("Error in get process items : no items found ")
	}
	return result, nil
}
