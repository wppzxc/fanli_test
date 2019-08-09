package utils

import (
	"fmt"
	"github.com/lxn/win"
	"github.com/wpp/fanli_test/pkg/types"
	"syscall"
	"testing"
	"time"
)

var (
	user32           = syscall.MustLoadDLL("user32")
	openClipboard    = user32.MustFindProc("OpenClipboard")
	closeClipboard   = user32.MustFindProc("CloseClipboard")
	emptyClipboard   = user32.MustFindProc("EmptyClipboard")
	getClipboardData = user32.MustFindProc("GetClipboardData")
	setClipboardData = user32.MustFindProc("SetClipboardData")
	loadImage        = user32.MustFindProc("LoadImageW")
	
	kernel32     = syscall.NewLazyDLL("kernel32")
	globalAlloc  = kernel32.NewProc("GlobalAlloc")
	globalFree   = kernel32.NewProc("GlobalFree")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	lstrcpy      = kernel32.NewProc("lstrcpyW")
)



func Test02(t *testing.T) {
	oldItems := []types.Item{{
		Id:             2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2582,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2585,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}}
	newItems := []types.Item{{
		Id:             2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2583,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2582,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2585,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}}
	
	result := GetDiffItems(oldItems, newItems)
	fmt.Println(result)
}

func Test03(t *testing.T) {
	for {
		p, err := syscall.UTF16PtrFromString("bigben")
		if err != nil {
			fmt.Println(err)
		}
		h2 := win.FindWindow(nil, p)
		re := SetForegroundWindow(h2)
		fmt.Println(re)
	}
}

func TestGetDiffItems(t *testing.T) {
	oldItems := []types.Item{{
		Id:             2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2584,
		ExtendDocument: "####拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}}
	newItems := []types.Item{{
		Id:             2583,
		ExtendDocument: "##拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2582,
		ExtendDocument: "#拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2584,
		ExtendDocument: "####拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}, {
		Id:             2581,
		ExtendDocument: "###拼多多免单来袭请注意看清要求\n###1.领取【1】元券，拍【零痕补水保湿面膜2片】选项，券后【2.8】元拍下\n###2.实发【抽纸一包】\n###3.禁止使用平台券，禁止联系商家咨询\n###4.有任何问题找群主\n###5.券无代表活动结束\n###6.拍完付款后请重新进入活动网址，点订单页面查询是否正常",
		GoodsImageUrl:  "",
		GoodsName:      "零痕天丝面膜补水保湿收缩毛孔美白淡斑玻尿酸精华淡化痘印男女",
		RefundAmount:   "2.80",
		MinGroupPrice:  "3.80",
		CouponDiscount: "1.00",
		StartTime:      1563197400,
		StopTime:       1563283800,
	}}
	result := newItems
	for _, o := range oldItems {
		for i, n := range newItems {
			if o.Id == n.Id {
				result = append(result[:i], result[i+1:]...)
				break
			}
		}
	}
	fmt.Println(result)
}

func TestSetClipboardImage(t *testing.T) {
	location, err := syscall.UTF16PtrFromString("C:\\Users\\wupengpeng\\Desktop\\kube-scheduler.jpg")
	if err != nil {
		fmt.Println(err)
	}
	imgH := win.LoadImage(0,
		location,
		win.IMAGE_BITMAP, 0, 0,
		win.LR_LOADFROMFILE)
	win.OpenClipboard(0)
	win.EmptyClipboard()
	win.SetClipboardData(win.CF_BITMAP, imgH)
	win.CloseClipboard()
}

func TestSetClipboardImage02(t *testing.T) {
	err := waitOpenClipboard()
	if err != nil {
		fmt.Println(err)
	}
	defer closeClipboard.Call()
	
	r, _, err := emptyClipboard.Call(0)
	if r == 0 {
		fmt.Println(err)
	}
	location, err := syscall.UTF16PtrFromString("C:\\Users\\wupengpeng\\Desktop\\kube-scheduler.jpg")
	if err != nil {
		fmt.Println(err)
	}
	
	imgH, _, err := loadImage.Call(0,
		uintptr(*location),
		win.IMAGE_BITMAP, 0, 0,
		win.LR_LOADFROMFILE)
	
	r, _, err = setClipboardData.Call(win.CF_BITMAP, imgH)
	if r == 0 {
		fmt.Println(err)
	}
	
}

func waitOpenClipboard() error {
	started := time.Now()
	limit := started.Add(time.Second)
	var r uintptr
	var err error
	for time.Now().Before(limit) {
		r, _, err = openClipboard.Call(0)
		if r != 0 {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return err
}
